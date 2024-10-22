# Modelo para los datos del cóctel
class Cocktail(BaseModel):
    idDrink: str
    strDrink: str
    strCategory: str
    strAlcoholic: str
    strGlass: str
    strInstructions: str
    strDrinkThumb: str