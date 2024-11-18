#%% extract text from pdfs

import os
from pypdf import PdfReader

def extract_text_from_pdf(file_path):
    pdf = PdfReader(file_path)
    text = ""
    for page in pdf.pages:
        text += page.extract_text()
    return text

# Lista de archivos PDF
pdf_files = ["./estatutoDeLosTrabajadores.pdf", "./prevencionDeRiesgosLaborales.pdf"]
documents = [extract_text_from_pdf(file) for file in pdf_files]

#%% create chunks of data which we will embed
MODEL_NAME = "all-MiniLM-L6-v2"

from sentence_transformers import SentenceTransformer
from langchain.text_splitter import RecursiveCharacterTextSplitter

# Initialize the model
model = SentenceTransformer(MODEL_NAME)

# Split the document text into smaller chunks
text_splitter = RecursiveCharacterTextSplitter(chunk_size=256, chunk_overlap=50)

all_chunks = []
all_embeddings = []
all_metadatas = []
all_ids = []

for i, doc in enumerate(documents):
    chunks = text_splitter.split_text(doc)
    embeddings = model.encode(chunks)
    
    all_chunks.extend(chunks)
    all_embeddings.extend(embeddings)
    all_metadatas.extend([{"source": pdf_files[i], "chunk": nchunk} for nchunk in range(len(chunks))])
    all_ids.extend([f"doc_{i}_chunk_{nchunk}" for nchunk in range(len(chunks))])

#%% create chromaDB
import chromadb
chroma_client = chromadb.PersistentClient(path="./chromadb")
collection = chroma_client.get_or_create_collection(name="pdf_embeddings")

#%% create collection or just get it

if collection.count() == 0:
    collection.upsert(
        documents=all_chunks,
        embeddings=all_embeddings,
        metadatas=all_metadatas,
        ids=all_ids
    )


#%% tests
from sentence_transformers import SentenceTransformer

model = SentenceTransformer(MODEL_NAME)

def test_query(query):
    query_embedding = model.encode(query)
    results = collection.query(
        query_embeddings=[query_embedding],
        n_results=5
    )
    print(f"Query: {query}\nResultado: {results}\n")

    
# Test 1: Pregunta sobre derechos laborales
test_query("¿Cuántos días de vacaciones tiene derecho un trabajador?")

# Test 2: Pregunta sobre seguridad en el trabajo
test_query("¿Qué formación en prevención de riesgos laborales debe recibir el trabajador?")

test_query("¿Qué me puedes decir de los contratos fijo discontinuos?")

test_query("¿Qué me puedes decir del artículo 16 del estatuto de los trabajadores?")

test_query("""¿El calendario de vacaciones se fijará en cada empresa. El
trabajador conocerá las fechas que le correspondan dos
meses antes, al menos, del comienzo del disfrute.?""")


# %% LLM ingest

from transformers import AutoModelForCausalLM, AutoTokenizer
import torch

# Configura el modelo LLM (por ejemplo, Llama 2)
model_name = "tiiuae/falcon-7b-instruct"  # meta-llama/Llama-2-7b-hf
tokenizer = AutoTokenizer.from_pretrained(model_name)
model = AutoModelForCausalLM.from_pretrained(model_name, torch_dtype=torch.float16, device_map="auto")

def query_with_llm(query):
    # Paso 1: Buscar contexto relevante en ChromaDB
    query_embedding = model.encode(query)
    results = collection.query(
        query_embeddings=[query_embedding],
        n_results=5
    )
    
    # Paso 2: Construir el contexto a partir de los resultados
    contexto = "\n".join([doc for doc in results['documents']])
    prompt = f"Contexto: {contexto}\nPregunta: {query}\nRespuesta:"
    
    # Paso 3: Llamar al LLM con el prompt construido
    inputs = tokenizer(prompt, return_tensors="pt").to("cuda")
    with torch.no_grad():
        outputs = model.generate(**inputs, max_new_tokens=150)

    # Decodificar y mostrar la respuesta
    respuesta = tokenizer.decode(outputs[0], skip_special_tokens=True)
    print(f"Pregunta: {query}\nRespuesta del LLM: {respuesta}\n")

# Prueba con una consulta
query_with_llm("¿Cuántos días de vacaciones tiene derecho un trabajador?")


# %%
