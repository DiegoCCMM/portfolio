import uvicorn
from fastapi_app import app  # Importar la instancia de FastAPI desde fastapi_app.py

if __name__ == "__main__":
    uvicorn.run(app, host="127.0.0.1", port=8001, log_level="info")

