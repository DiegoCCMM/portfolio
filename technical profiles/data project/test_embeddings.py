#%% imports
import gc
import torch
import chromadb
import logging
from tqdm import tqdm
from sentence_transformers import SentenceTransformer

#%% Configuración del logging
logging.basicConfig(
    filename="test_embeddings.log",
    level=logging.INFO,
    format="%(asctime)s - %(levelname)s - %(message)s"
)

#%% Inicializar ChromaDB en modo de solo lectura
chroma_client = chromadb.PersistentClient(path="./chromadb_storage")
collections = chroma_client.list_collections()
logging.info(f"Se encontraron {len(collections)} colecciones en ChromaDB.")

#%% Función para obtener el modelo de embeddings a partir del nombre de la colección
def get_model_from_collection_name(collection_name):
    model_name = collection_name.split("_size")[0]
    return SentenceTransformer(model_name)

#%% Definir las queries para probar
queries = [
    "¿Cuántos días de vacaciones tiene derecho un trabajador?",
    "¿Cuáles son los derechos de los trabajadores menores de edad?",
    "¿Qué debe incluir un contrato laboral?",
    "Normas de prevención de riesgos laborales.",
    "Derechos de conciliación para trabajadores fijos-discontinuos."
]

#%% Función para probar una colección con todas las queries
def test_collection(collection, model):
    collection_name = collection.name
    logging.info(f"\nProbando la colección: {collection_name}")

    for query in tqdm(queries, desc=f"Queries para {collection_name}"):
        try:
            # Generar el embedding de la query
            query_embedding = model.encode(query, convert_to_numpy=True)

            # Realizar la búsqueda usando el embedding de la query
            results = collection.query(
                query_embeddings=[query_embedding],
                n_results=5,
                include=["documents", "metadatas", "distances"]
            )

            # Registrar los resultados en el log
            logging.info(f"\nQuery: {query}")
            for i, doc in enumerate(results["documents"][0]):
                metadata = results["metadatas"][0][i]
                distance = results["distances"][0][i]
                source = metadata.get("source", "desconocido")

                logging.info(f"- Resultado {i+1}: (Distancia: {distance:.4f})")
                logging.info(f"  Documento: {doc[:100]}...")  # Mostrar solo los primeros 100 caracteres
                logging.info(f"  Fuente: {source}")

        except Exception as e:
            logging.error(f"Error al procesar la query '{query}' para la colección '{collection_name}': {e}")

#%% Probar todas las colecciones encontradas
for collection in collections:
    test_collection(collection, get_model_from_collection_name(collection.name))

    torch.cuda.empty_cache()
    gc.collect()

logging.info("\nPruebas completadas para todas las colecciones.")
