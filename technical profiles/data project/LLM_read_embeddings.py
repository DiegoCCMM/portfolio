#%%
import torch
from transformers import GPT2LMHeadModel, GPT2Tokenizer
from sentence_transformers import SentenceTransformer

device = "cuda" if torch.cuda.is_available() else "cpu"
model_name = "gpt2"

# Cargar el modelo y tokenizer
tokenizer = GPT2Tokenizer.from_pretrained(model_name)
# Configurar el pad_token_id
tokenizer.pad_token = tokenizer.eos_token
tokenizer.add_special_tokens({'pad_token': '[PAD]'})
model = GPT2LMHeadModel.from_pretrained(model_name).to(device)
model.resize_token_embeddings(len(tokenizer))
model.eval()


#%%
def paraphrase_text(text):
    # Tokenizar la entrada y agregar atención
    inputs = tokenizer(
        text,
        return_tensors="pt",
        padding=True,
        truncation=True
    ).to(device)

    with torch.no_grad():
        # Generar texto con el modelo
        output_ids = model.generate(
            inputs["input_ids"],
            attention_mask=inputs["attention_mask"],
            max_length=500,
            num_return_sequences=1,
            temperature=0.7,
            top_k=50,
            top_p=0.9,
            do_sample=True
        )
    
    # Decodificar el texto generado
    paraphrased_text = tokenizer.decode(output_ids[0], skip_special_tokens=True)
    return paraphrased_text

#%%
import chromadb

# Inicializar el cliente ChromaDB en modo de solo lectura
chroma_client = chromadb.PersistentClient(path="./chromadb_storage")
collections = chroma_client.list_collections()

# Definir una batería de queries para probar
queries = [
    "¿Cuántos días de vacaciones tiene derecho un trabajador?",
    "¿Cuáles son los derechos de los trabajadores menores de edad?",
    "¿Qué debe incluir un contrato laboral?",
    "Normas de prevención de riesgos laborales.",
    "Derechos de conciliación para trabajadores fijos-discontinuos."
]

def rag_pipeline(query, collection):
    # Obtener el embedding del query usando un modelo de embeddings
    model_name = collection.name.split("_size")[0]
    embedding_model = SentenceTransformer(model_name)
    query_embedding = embedding_model.encode(query, convert_to_numpy=True)

    # Buscar en la colección de ChromaDB
    try:
        search_results = collection.query(
            query_embeddings=[query_embedding],
            n_results=3,
            include=["documents", "metadatas"]
        )

        # Concatenar los documentos encontrados
        combined_text = " ".join(search_results["documents"][0])

        # Reescribir o parafrasear el texto combinado
        paraphrased_output = paraphrase_text(combined_text)
        return paraphrased_output

    except Exception as e:
        print(f"Error al procesar la query '{query}': {e}")
        return None

#%%
import gc

for collection in collections:
    print(f"\nProbando la colección: {collection.name}")
    for query in queries:
        response = rag_pipeline(query, collection)
        if response:
            print(f"\nQuery: {query}")
            print(f"Respuesta Generada: {response}\n")
        else:
            print(f"No se pudo generar una respuesta para la query: {query}")
    
    # Liberar memoria
    torch.cuda.empty_cache()
    gc.collect()

print("Pruebas completadas para todas las colecciones.")

# %%
