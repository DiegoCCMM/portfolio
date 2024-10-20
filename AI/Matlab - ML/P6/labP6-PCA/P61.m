%% Practica 6.1: PCA 

clear all
close all

% Leer los datos originales en la variable X
load P61

% Graficar los datos originales
figure(1);
axis equal;
grid on;
hold on;
plot3(X(:,1),X(:,2),X(:,3),'b.');
xlabel ('X');
ylabel ('Y');
zlabel ('Z');
pause

% Estandarizar los datos (solo hace falta centrarlos)

% Graficar los datos centrados

% Calcular la matrix de covarianza muestral de los datos centrados

% Aplicar PCA para obtener los vectores propios y valores propios

% Ordenar los vectores y valores preprios de mayor a menor valor propio

% Graficar en color rojo cada vector propio * 3 veces la raiz de su 
% correspondiente valor propio

% Graficar la variabilidad que se mantiene si utilizas los tres primeros
% vectores propios, los dos primeros, o solo el primer vector propio

% Aplicar PCA para reducir las dimensiones de los datos y mantener al menos
% el 90% de la variabilidad

% Graficar aparte los datos z proyectados según el resultado anterior

% Graficar en verde los datos reproyectados \hat{x} en la figura original
