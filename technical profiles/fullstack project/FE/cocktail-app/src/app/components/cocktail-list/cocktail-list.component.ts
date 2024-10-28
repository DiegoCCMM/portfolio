import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CocktailService, Cocktail } from '../../services/cocktail.service';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-cocktail-list',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './cocktail-list.component.html',
  styleUrls: ['./cocktail-list.component.css'],
  providers: [CocktailService],
})

export class CocktailListComponent {
  cocktails$: Observable<Cocktail[]>;

  constructor(private cocktailService: CocktailService) {
    this.cocktails$ = this.cocktailService.getCocktails();
    console.log("FE: recibido cocktails$ ", this.cocktails$)
  }
}
