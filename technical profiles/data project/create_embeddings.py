#%% IMPORTS
import os
import gc
import torch
from pypdf import PdfReader
import chromadb
from sentence_transformers import SentenceTransformer
from langchain.text_splitter import RecursiveCharacterTextSplitter
from tqdm import tqdm

#%% INIT DB AND DATA
# Inicializar cliente de ChromaDB con persistencia
chroma_client = chromadb.PersistentClient(path="./chromadb_storage")

# Función para extraer texto de un archivo PDF
def extract_text_from_pdf(file_path):
    pdf = PdfReader(file_path)
    text = ""
    for page in pdf.pages:
        text += page.extract_text()
    return text

# Lista de archivos PDF
pdf_files = ["./estatutoDeLosTrabajadores.pdf", "./prevencionDeRiesgosLaborales.pdf"]
documents = [extract_text_from_pdf(file) for file in pdf_files]

#%% STUDY OVER THIS
# Modelos para probar
models = [
    "all-MiniLM-L6-v2",
    "all-MiniLM-L12-v2",
    "all-distilroberta-v1",
    "paraphrase-MiniLM-L6-v2",
    "multi-qa-mpnet-base-dot-v1",
    "all-mpnet-base-v2",
    "sentence-t5-base"
]

# Chunk sizes y overlaps para probar
chunk_sizes = [128, 256, 512]
chunk_overlaps = [30, 50, 100]

#%% run creation of collection
# Función para procesar e insertar documentos en ChromaDB
def process_and_insert(model_name, chunk_size, chunk_overlap, documents):
    # Nombre de la colección basado en el modelo, size y overlap
    collection_name = f"{model_name}_size{chunk_size}_overlap{chunk_overlap}"
    collection = chroma_client.get_or_create_collection(name=collection_name)

    if collection.count() != 0:
        print(f"Datos ya presentes en la colección: {collection_name}. Inserción omitida.")
    else:
        print(f"Iniciando calculos para colección: {collection_name}")

        model = SentenceTransformer(model_name)
        text_splitter = RecursiveCharacterTextSplitter(chunk_size=chunk_size, chunk_overlap=chunk_overlap)

        all_chunks = []
        all_embeddings = []
        all_metadatas = []
        all_ids = []


        for doc_idx, doc in enumerate(documents):

            chunks = text_splitter.split_text(doc)
            embeddings = model.encode(chunks, convert_to_numpy=True)
            all_chunks.extend(chunks)
            all_embeddings.extend(embeddings)
            all_metadatas.extend([{"source": pdf_files[doc_idx], "chunk": i} for i in range(len(chunks))])
            all_ids.extend([f"doc_{doc_idx}_chunk_{i}" for i in range(len(chunks))])

    
    # Insertar datos si la colección está vacía
    
        print(f"Iniciando inserción para colección: {collection_name}")
        collection.upsert(
            documents=all_chunks,
            embeddings=all_embeddings,
            metadatas=all_metadatas,
            ids=all_ids
        )
        print(f"Datos insertados para colección: {collection_name}")
    
    # Liberar heap SIEMPRE (torch reserva memoria y no la usa)

    if all(var in locals() for var in ['all_chunks', 'all_embeddings', 'all_metadatas', 'all_ids']): 
        del all_chunks, all_embeddings, all_metadatas, all_ids    
    torch.cuda.empty_cache()
    gc.collect()
        

# Ejecutar pruebas para cada combinación de modelo, chunk_size y chunk_overlap
total_iterations = len(models) * len(chunk_sizes) * len(chunk_overlaps)
with tqdm(total=total_iterations, desc="Procesando configuraciones") as pbar:
    for model_name in models:
        for chunk_size in chunk_sizes:
            for chunk_overlap in chunk_overlaps:
                process_and_insert(model_name, chunk_size, chunk_overlap, documents)
                pbar.update(1)
print("Todas las configuraciones procesadas e insertadas en ChromaDB.")

# %%
