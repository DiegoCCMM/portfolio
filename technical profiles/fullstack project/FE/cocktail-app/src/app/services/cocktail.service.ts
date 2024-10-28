import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Cocktail {
  strDrink: string;
  strAlcoholic: string;
  strDrinkThumb: string;
  strInstructions: string;
}

@Injectable({
  providedIn: 'root'
})
export class CocktailService {
  private apiUrl = 'http://localhost:8001';

  constructor(private http: HttpClient) {}

  getCocktails(): Observable<Cocktail[]> {
    return this.http.get<Cocktail[]>(`${this.apiUrl}/cocktails/`);
  }
}
