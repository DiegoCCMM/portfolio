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

%% Estandarizar los datos (solo hace falta centrarlos)
[Xi, mu] = estandarizar(X);
%% Graficar los datos centrados
figure(2);
axis equal;
grid on;
hold on;
plot3(Xi(:,1),Xi(:,2),Xi(:,3),'b.');
xlabel ('X');
ylabel ('Y');
zlabel ('Z');
%% Calcular la matrix de covarianza muestral de los datos centrados
CovMues = cov(X);
%% Aplicar PCA para obtener los vectores propios y valores propios
[U, lambda] = eig(CovMues);
%% Ordenar los vectores y valores propios de mayor a menor valor propio
[U,lambda] = ordenacionVVP(U,lambda);
%% Graficar en color rojo cada vector propio * 3 veces la raiz de su correspondiente valor propio
figure(3);
for i=1:size(U)
    Uplot(:,i)=U(:,i).*((sqrt(max(lambda(:,i))))*3);
end
hold on
plot3([0,Uplot(1,1)],[0,Uplot(2,1)], [0,Uplot(3,1)], 'r-');
plot3([0,Uplot(1,2)],[0,Uplot(2,2)], [0,Uplot(3,2)], 'r-');
plot3([0,Uplot(1,3)],[0,Uplot(2,3)], [0,Uplot(3,3)], 'r-');
axis equal;
grid on;
hold on;
plot3(Xi(:,1),Xi(:,2),Xi(:,3),'b.');
hold off

%% Graficar la variabilidad que se mantiene si utilizas los dos primeros, o solo el primer vector propio
figure(4);  %dos primeros vectores
for i=1:size(U)
    Uplot(:,i)=U(:,i).*((sqrt(max(lambda(:,i))))*3);
end
hold on
plot3([0,Uplot(1,1)],[0,Uplot(2,1)], [0,Uplot(3,1)], 'r-');
plot3([0,Uplot(1,2)],[0,Uplot(2,2)], [0,Uplot(3,2)], 'r-');
axis equal;
grid on;
hold on;
plot3(Xi(:,1),Xi(:,2),Xi(:,3),'b.');
hold off
%% primer vector propio
figure(5);  % primero
for i=1:size(U)
    Uplot(:,i)=U(:,i).*((sqrt(max(lambda(:,i))))*3);
end
hold on
plot3([0,Uplot(1,1)],[0,Uplot(2,1)], [0,Uplot(3,1)], 'r-');
axis equal;
grid on;
hold on;
plot3(Xi(:,1),Xi(:,2),Xi(:,3),'b.');
hold off
%% Aplicar PCA para reducir las dimensiones de los datos y mantener al menos el 90% de la variabilidad
[k]=elegirK(lambda, 90);
%% Graficar aparte los datos z proyectados seg�n el resultado anterior
Ut=transpose(U(:,1:k));
Xit=transpose(Xi);
Zt=ones(2,1000);
%% calculo z
for i=1:1000
        Zt(:,i)=Ut*Xit(:,i);
end
%% graficar z

Z=transpose(Zt);

figure(6);
axis equal;
grid on;
hold on;
plot(Z(:,1),Z(:,2),'k.');
xlabel ('X');
ylabel ('Y');
zlabel ('Z');
%% Graficar en verde los datos reproyectados \hat{x} en la figura original
U=transpose(Ut);
Xgorrot=U*Zt;
Xgorro=transpose(Xgorrot);
figure(7);
axis equal;
grid on;
hold on;
plot3(Xgorro(:,1),Xgorro(:,2),Xgorro(:,3),'b.');
xlabel ('X');
ylabel ('Y');
zlabel ('Z');
