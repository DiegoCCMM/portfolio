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

%% SVD
[U,S,V] = svd(X);
Vt=transpose(V);
%% Graficar las primeras 5 componentes
for k = 1:5
    figure(2);
    imshow(U(:,k)*S(k,k)*Vt(k,:));
    colormap(gray);
    axis off;
end

%% Graficar la reconstrucci�n con las primeras 1, 2, 5, 10, 20, y total
% de componentes
for k = [1 2 5 10 20 rank(X)]
    figure(3);
    imshow(U(:,1:k)*S(1:k,1:k)*Vt(1:k,:));
    colormap(gray);
    axis off;
end
%% Graficar la reconsrtucci�n con las primeras k componentes
[k]=elegirK(S, 90);
figure(4);
imshow(U(:,1:k)*S(1:k,1:k)*Vt(1:k,:));
colormap(gray);
axis off;
%% Calcular y mostrar el ahorro en espacio
[m,n]=size(X);
R=(m*n)./ (k*(m+n))

figure(3);
plot((m*n)./ ((1:1:k)*(m+n)));
xlabel ('K');
ylabel ('R');

