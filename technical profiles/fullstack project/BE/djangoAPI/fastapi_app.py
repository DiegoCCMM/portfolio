from fastapi import FastAPI
from motor.motor_asyncio import AsyncIOMotorClient
from pydantic import BaseModel
from cocktail_model import Cocktail

# Configurar FastAPI
app = FastAPI()

# Conectar a la base de datos MongoDB
client = AsyncIOMotorClient('mongodb://localhost:27017')
db = client.cocktail_db
cocktails_collection = db.cocktails

# Obtener todos los cócteles
@app.get("/cocktails/", response_model=list[Cocktail])
async def get_cocktails():
    cocktails = await cocktails_collection.find().to_list(1000)
    if not cocktails:
        raise HTTPException(status_code=404, detail="No cocktails found")
    return cocktails

# Buscar cóctel por nombre
@app.get("/cocktails/{name}", response_model=list[Cocktail])
async def search_cocktail(name: str):
    cocktails = await cocktails_collection.find({"strDrink": {"$regex": name, "$options": "i"}}).to_list(1000)
    if not cocktails:
        raise HTTPException(status_code=404, detail="No cocktails found")
    return cocktails