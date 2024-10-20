%% cargamos los datos
clear ; close all;
addpath(genpath('../minfunc'));

% Carga los datos y los permuta aleatoriamente
load('MNISTdata2.mat'); % Lee los datos: X, y, Xtest, ytest
rand('state',0);
p = randperm(length(y));
X = X(p,:);
y = y(p);

% apartado 3 % 

%% encontramos los modelos gaussianos
[best_lambda, listError_train, listError_cv] = KfoldCrossValidationLambda(5,X,y,1);

%% Re-entrena con todos los datos
modelo = entrenarGaussianas( X, y, 10, 1, best_lambda );
%% Predicción y clasificación
yCla = clasificacionBayesiana(modelo, 10, Xtest);
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

% apartado 4 % 

%% encontramos los modelos gaussianos
[best_lambda, listError_train, listError_cv] = KfoldCrossValidationLambda(5,X,y,0);

 %% Re-entrena con todos los datos
modelo = entrenarGaussianas( X, y, 10, 0, best_lambda );
%% Predicción y clasificación
yCla = clasificacionBayesiana(modelo, 10, Xtest);
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