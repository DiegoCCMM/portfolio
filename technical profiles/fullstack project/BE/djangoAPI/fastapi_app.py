from fastapi import FastAPI, Query
from fastapi.middleware.cors import CORSMiddleware
from motor.motor_asyncio import AsyncIOMotorClient
from pydantic import BaseModel
from cocktail_model import Cocktail

# Configurar FastAPI
app = FastAPI()

origins = [
    "http://localhost:4200",  # Cambia esto al puerto donde se ejecuta tu app Angular
    # Puedes añadir más orígenes si es necesario
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,  # Orígenes permitidos
    allow_credentials=True,
    allow_methods=["*"],  # Permitir cualquier método
    allow_headers=["*"],  # Permitir cualquier encabezado
)

# Conectar a la base de datos MongoDB
client = AsyncIOMotorClient('mongodb://localhost:27017')
db = client.cocktail_db
cocktails_collection = db.cocktails

# Obtener todos los cócteles
@app.get("/cocktails/", response_model=list[Cocktail])
async def get_cocktails():
    print("BE: Recibido request cocktails")
    cocktails = await cocktails_collection.find().to_list(1000)
    if not cocktails:
        raise HTTPException(status_code=404, detail="No cocktails found")
    return cocktails


# Buscar por campos
@app.get("/cocktails/search", response_model=list[Cocktail])
async def get_cocktails(name: str = Query(...), alcoholic: bool = Query(...)):
    # Definir el filtro basado en el parámetro alcoholic
    alcoholic_value = "Alcoholic" if alcoholic else "Non_Alcoholic"
    
    cocktails = await cocktails_collection.find(
        {"$and": [
            { "strAlcoholic": alcoholic_value },
            { "strDrink": name }
        ]
        }).to_list(1000)
    
    if not cocktails:
        raise HTTPException(status_code=404, detail="No cocktails found")
    
    return cocktails