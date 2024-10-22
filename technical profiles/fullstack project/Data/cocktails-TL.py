# %% connection
from pymongo import MongoClient
import json

# Conectar a MongoDB (puede ser una conexión remota o local)
client = MongoClient('mongodb://localhost:27017/')
db = client['cocktail_db']  # Nombre de la base de datos
cocktails_collection = db['cocktails']  # Nombre de la colección

# %% T
# Cargar los datos del archivo JSON
with open('cocktails.json', 'r') as f:
    cocktails_data = json.load(f)

# %% L
# Insertar todos los cócteles en la colección
if isinstance(cocktails_data, list):
    cocktails_collection.insert_many(cocktails_data)
else:
    print("Los datos no tienen el formato adecuado.")
print(f"Total de cócteles en la base de datos: {cocktails_collection.count_documents({})}")

# %%
