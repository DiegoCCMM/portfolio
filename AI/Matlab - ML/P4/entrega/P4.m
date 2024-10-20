%% Limpiamos

clear ; close all;

%% Practica 4 
% Based on exercise 3 of Machine Learning Online Class by Andrew Ng 
%

clear ; close all;
addpath(genpath('../minfunc'));

% Carga los datos y los permuta aleatoriamente
load('MNISTdata2.mat'); % Lee los datos: X, y, Xtest, ytest
rand('state',0);
p = randperm(length(y));
X = X(p,:);
y = y(p);

X = [ones(length(y),1) X];
Xtest = [ones(length(ytest),1) Xtest];
options.display = 'final'; %otros: 'iter' , 'none‘
options.method = 'pnewton0'; %por defecto: 'lbfgs'
options.useMex = 1;

%% crear partición y expansión

[ Xcv, ycv, Xtr, ytr] = particion(1, 5, X, y);
%% calculo de la lambda para ajustar
[best_size, listError_train, listError_cv] = KfoldCrossValidation(5,X,y);
%% calculo de theta gorda reentrenando
THETA=[];
theta0 = zeros(size(X,2),1);
for j=1:10
    yj = y == j;    

    theta = minFunc(@CosteLogReg, theta0, options, X, yj, best_size);

    THETA = [THETA, theta];
end
%% calculo de la predicción
        htr = 1./(1+exp(-(Xtr*THETA))); %LA hipotesis CON LA MEJOR THETA sobre todos los datos
        yhtr = zeros(size(htr,1),1);
        for i=1:size(htr,1)
            [val, loc] = max(htr(i,:));
            yhtr(i) = loc;
        end
        
        hcv = 1./(1+exp(-(Xcv*THETA))); %LA hipotesis CON LA MEJOR THETA sobre todos los datos
        yhcv = zeros(size(hcv,1),1);
        for i=1:size(hcv,1)
            [val, loc] = max(hcv(i,:));
            yhcv(i) = loc;
        end
        
        errorescv=yhcv~=ycv;

        errorestr=yhtr~=ytr;
        
        Tetr=sum(errorestr)/size(yhtr,1)  %error del train

        Tecv=sum(errorescv)/size(yhcv,1) %error del test
%% calculo de la matriz de confusión

matrizDeConfusion = zeros(11,11);
%encontramos los aciertos y los fallos
TP = find(yhcv==ycv);
fallos = find(yhcv ~= ycv);

for i=1:size(TP)
%sabiendo que son aciertos, sumamos en esa diagonal +1   
    matrizDeConfusion(yhcv(TP(i,1)),yhcv(TP(i,1))) = matrizDeConfusion(yhcv(TP(i,1)),yhcv(TP(i,1))) + 1;
end
for i=1:size(fallos)
%sabiendo que son fallos, sumamos el error, en la fila correcta, la columna
%decidida
    matrizDeConfusion(ycv(fallos(i,1)),yhcv(fallos(i,1))) = matrizDeConfusion(ycv(fallos(i,1)),yhcv(fallos(i,1))) + 1;
end
suma = sum(matrizDeConfusion);
for i=1:10
   matrizDeConfusion(i, 11) = sum(matrizDeConfusion(i,:));
   matrizDeConfusion(11,i)=suma(i);
end
suma = sum(matrizDeConfusion(:,11));
matrizDeConfusion(11,11)=suma;

matrizDeConfusion
%% recall,f1,precision
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
        
%% mostrar confusiones
verConfusiones(Xcv, ycv, yhcv);