%% Practica 6.2: PCA 

clear all
close all

% Leer la imagen
I = imread('turing.png');

% Convertirla a blanco y negro
BW = rgb2gray(I);

% Convertir los datos a double
X=im2double(BW);

% graficar la imagen
figure(1);
colormap(gray);
imshow(X);
axis off;
pause

% Aplicar PCA

% Graficar las primeras 5 componentes
for k = 1:5,
    figure(2);
    %imshow(Xhat);
    colormap(gray);
    axis off;
    pause
end

% Graficar la reconstrucción con las primeras 1, 2, 5, 10, 20, y total
% de componentes
for k = [1 2 5 10 20 rank(X)],
    figure(3);
    %imshow(Xhat);
    colormap(gray);
    axis off;
    pause
end

% Encontrar el valor de k que mantenga al menos el 90% de la variabilidad

% Graficar la reconsrtucción con las primeras k componentes

% Calcular y mostrar el ahorro en espacio
