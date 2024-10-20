% load images 
% images size is 20x20. 
clear
close all

load('MNISTdata2.mat'); 

nrows=20;
ncols=20;

nimages = size(X,1);

%% Show the images
for I=1:40:nimages 
    imshow(reshape(X(I,:),nrows,ncols))
    pause(0.1)
end


%% Estandarizar los datos (solo hace falta centrarlos)
[Xi, mu] = estandarizar(X);
%% Calcular la matrix de covarianza muestral de los datos centrados
CovMues = cov(X);
%% Aplicar PCA para obtener los vectores propios y valores propios
[U, lambda] = eig(CovMues);
%% Ordenar los vectores y valores propios de mayor a menor valor propio
[U,lambda] = ordenacionVVP(U,lambda);
%% Aplicar PCA para reducir las dimensiones de los datos y mantener al menos el 90% de la variabilidad
[k]=elegirK(lambda, 90);
%% Graficar aparte los datos z proyectados sobre los dos más importantes
Ut=transpose(U(:,1:k));
Xit=transpose(Xi);
Zt=ones(52,400);
%% calculo z
for i=1:4000
        Zt(:,i)=Ut*Xit(:,i);
end

%% z should contain the projections over the first two PC
% now is just a random matrix
z=transpose(Zt);


% Muestra las dos componentes principales
figure(100)
clf, hold on
plotwithcolor(z(:,1:2), y);

%% Calculamos todos los datos proyectados
U=transpose(Ut);
Xgorrot=U*transpose(z);
Xgorro=transpose(Xgorrot);

%% Use classifier from previous labs on the projected space
%para este caso, elegiremos como clases sencillas el 1 y el 10 y como clases complicadas
%el 4 y el 8.
yIndex=find( y==1 | y==10);
yIndexM= find( ytest==1 | ytest==10);
zSelectedB=z(yIndex,:);
ySelectedB=y(yIndex);
ytestB=ytest(yIndexM);
xtestB=Xtest(yIndexM,:);

yIndex= find( y==4 | y==8);
yIndexM= find( ytest==4 | ytest==8);
zSelectedM=z(yIndex,:);
ySelectedM=y(yIndex);
ytestM=ytest(yIndexM);
xtestM=Xtest(yIndexM,:);
%% proyectamos los datos seleccionados
U=transpose(Ut);
XgorrotB=U*transpose(zSelectedB);
XgorroB=transpose(XgorrotB);
clases=[1,10];

%% clasificador de las buenas ******************************************************************************************************

[best_lambda, listError_train, listError_cv] = KfoldCrossValidationLambda(5,XgorroB,ySelectedB,1,clases);


%% Re-entrena con todos los datos
modelo = entrenarGaussianas( XgorroB, ySelectedB, 10, 1, best_lambda);
%% Predicción y clasificación
yCla = clasificacionBayesiana(modelo, 10, xtestB, clases);
%% Calculo de matriz de confusión

matrizDeConfusion = zeros(11,11);
%encontramos los aciertos y los fallos
TP = find(yCla==ytestB);
fallos = find(yCla ~= ytestB);

for i=1:size(TP)
%sabiendo que son aciertos, sumamos en esa diagonal +1   
    matrizDeConfusion(yCla(TP(i,1)),yCla(TP(i,1))) = matrizDeConfusion(yCla(TP(i,1)),yCla(TP(i,1))) + 1;
end
for i=1:size(fallos)
%sabiendo que son fallos, sumamos el error, en la fila correcta, la columna
%decidida
    matrizDeConfusion(ySelectedB(fallos(i,1)),yCla(fallos(i,1))) = matrizDeConfusion(ySelectedB(fallos(i,1)),yCla(fallos(i,1))) + 1;
end
suma = sum(matrizDeConfusion);
for i=1:10
   matrizDeConfusion(i, 11) = sum(matrizDeConfusion(i,:));
   matrizDeConfusion(11,i)=suma(i);
end
suma = sum(matrizDeConfusion(:,11));
matrizDeConfusion(11,11)=suma;

matrizDeConfusion

%% calculo de f1, recall, precision y ver confusiones
recall=[];
precision=[];
for i=1:10
    recall = [recall,matrizDeConfusion(i,i)/matrizDeConfusion(i,11)];
    precision = [precision,matrizDeConfusion(i,i)/matrizDeConfusion(11,i)];
    recall = mean(recall);
    precision = mean(precision);
end
recall
precision
f1=(2*(recall*precision))/(recall+precision);
f1
verConfusiones(xtestB,ytestB,yCla);


%% proyectamos los datos seleccionados ***********************************************************************************
U=transpose(Ut);
XgorrotM=U*transpose(zSelectedM);
XgorroM=transpose(XgorrotM);
clases=[4,8];
%% clasificador MALO  *******************************************************************************************************+*
[best_lambda, listError_train, listError_cv] = KfoldCrossValidationLambda(5,XgorroM,ySelectedM,1,clases);


%% Re-entrena con todos los datos
modelo = entrenarGaussianas( XgorroM, ySelectedM, 10, 1, best_lambda);
%% Predicción y clasificación
yCla = clasificacionBayesiana(modelo, 10, xtestM, clases);
%% Calculo de matriz de confusión

matrizDeConfusion = zeros(11,11);
%encontramos los aciertos y los fallos
TP = find(yCla==ytestM);
fallos = find(yCla ~= ytestM);

for i=1:size(TP)
%sabiendo que son aciertos, sumamos en esa diagonal +1   
    matrizDeConfusion(yCla(TP(i,1)),yCla(TP(i,1))) = matrizDeConfusion(yCla(TP(i,1)),yCla(TP(i,1))) + 1;
end
for i=1:size(fallos)
%sabiendo que son fallos, sumamos el error, en la fila correcta, la columna
%decidida
    matrizDeConfusion(ySelectedM(fallos(i,1)),yCla(fallos(i,1))) = matrizDeConfusion(ySelectedM(fallos(i,1)),yCla(fallos(i,1))) + 1;
end
suma = sum(matrizDeConfusion);
for i=1:10
   matrizDeConfusion(i, 11) = sum(matrizDeConfusion(i,:));
   matrizDeConfusion(11,i)=suma(i);
end
suma = sum(matrizDeConfusion(:,11));
matrizDeConfusion(11,11)=suma;

matrizDeConfusion

%% calculo de f1, recall, precision y ver confusiones
recall=[];
precision=[];
for i=1:10
    recall = [recall,matrizDeConfusion(i,i)/matrizDeConfusion(i,11)];
    precision = [precision,matrizDeConfusion(i,i)/matrizDeConfusion(11,i)];
    recall = mean(recall);
    precision = mean(precision);
end
recall
precision
f1=(2*(recall*precision))/(recall+precision);
f1
verConfusiones(xtestB,ytestB,yCla);
%% clasificador TODOS LOS DATOS  *******************************************************************************************************+*
p = randperm(length(y));
Xgorro = Xgorro(p,:);
y = y(p);
clases=[1:10];
[best_lambda, listError_train, listError_cv] = KfoldCrossValidationLambda(5,Xgorro,y,1,clases);


%% Re-entrena con todos los datos
modelo = entrenarGaussianas( Xgorro, y, 10, 1, best_lambda);
%% Predicción y clasificación
yCla = clasificacionBayesiana(modelo, 10, Xtest, clases);
%% Calculo de matriz de confusión

matrizDeConfusion = zeros(11,11);
%encontramos los aciertos y los fallos
TP = find(yCla==ytest);
fallos = find(yCla ~= ytest);

for i=1:size(TP)
%sabiendo que son aciertos, sumamos en esa diagonal +1   
    matrizDeConfusion(yCla(TP(i,1)),yCla(TP(i,1))) = matrizDeConfusion(yCla(TP(i,1)),yCla(TP(i,1))) + 1;
end
for i=1:size(fallos)
%sabiendo que son fallos, sumamos el error, en la fila correcta, la columna
%decidida
    matrizDeConfusion(y(fallos(i,1)),yCla(fallos(i,1))) = matrizDeConfusion(y(fallos(i,1)),yCla(fallos(i,1))) + 1;
end
suma = sum(matrizDeConfusion);
for i=1:10
   matrizDeConfusion(i, 11) = sum(matrizDeConfusion(i,:));
   matrizDeConfusion(11,i)=suma(i);
end
suma = sum(matrizDeConfusion(:,11));
matrizDeConfusion(11,11)=suma;

matrizDeConfusion

%% calculo de f1, recall, precision y ver confusiones
recall=[];
precision=[];
for i=1:10
    recall = [recall,matrizDeConfusion(i,i)/matrizDeConfusion(i,11)];
    precision = [precision,matrizDeConfusion(i,i)/matrizDeConfusion(11,i)];
    recall = mean(recall);
    precision = mean(precision);
end
recall
precision
f1=(2*(recall*precision))/(recall+precision);
f1
verConfusiones(Xtest,ytest,yCla);



