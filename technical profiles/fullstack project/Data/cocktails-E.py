import requests

# URL base de la API
base_url = "https://www.thecocktaildb.com/api/json/v1/1/"

# Función para obtener todos los cócteles por primera letra
def get_cocktails_by_letter(letter):
    url = f"{base_url}search.php?f={letter}"
    response = requests.get(url)
    if response.status_code == 200:
        return response.json()
    else:
        print(f"Error {response.status_code} al obtener cócteles.")
        return None

# Extraer todos los cócteles
def extract_all_cocktails():
    cocktails = []
    for letter in 'abcdefghijklmnopqrstuvwxyz':
        data = get_cocktails_by_letter(letter)
        if data and 'drinks' in data and data['drinks'] is not None:
            cocktails.extend(data['drinks'])
    return cocktails

# Guardar la información de los cócteles
def save_cocktails_to_file(cocktails, filename='cocktails.json'):
    import json
    with open(filename, 'w') as f:
        json.dump(cocktails, f, indent=4)

# Ejecutar el script
if __name__ == "__main__":
    all_cocktails = extract_all_cocktails()
    if all_cocktails:
        save_cocktails_to_file(all_cocktails)
        print(f"Se han guardado {len(all_cocktails)} cócteles en 'cocktails.json'.")
    else:
        print("No se pudieron obtener cócteles.")

