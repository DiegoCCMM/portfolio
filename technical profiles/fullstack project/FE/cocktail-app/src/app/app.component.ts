import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CocktailListComponent } from './components/cocktail-list/cocktail-list.component';

@Component({
  selector: 'app-root',  // Asegúrate de que este selector coincida con <app-root> en index.html
  standalone: true,
  imports: [CommonModule, CocktailListComponent],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {}
