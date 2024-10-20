%% Based on exercise 2 of Machine Learning Online Class by Andrew Ng 

clear ; close all;

%% Load Data del primer apartado
%  The first two columns contains the exam scores and the third column
%  contains the label.

data = load('exam_data.txt');
y = data(:, 3);
N = length(y);
X = data(:, [1, 2]); 

plotData(X, y);
xlabel('Exam 1 score')
ylabel('Exam 2 score')
legend('Admitted', 'Not admitted')


%% Calcula una Solucion con costelogistico
X = [ones(N,1) X];
theta = transpose([1,1,1]);
[ Xcv, ycv, Xtr, ytr] = particion(2, 5, X, y);
options.display = 'final'; %otros: 'iter' , 'none‘
options.method = 'newton'; %por defecto: 'lbfgs'
theta = minFunc(@CosteLogistico, theta, options, Xtr, ytr);


%% sacamos la hipótesis y la utilizamos

htr = 1./(1+exp(-(Xtr*theta))); %LA hipotesis CON LA MEJOR THETA sobre todos los datos
yhtr = ones(size(htr));

hcv = 1./(1+exp(-(Xcv*theta)));
yhcv = ones(size(hcv));

for i=1:80
   if htr(i) < 0.5
       yhtr(i) = 0; %predicción para el train
   end
end

for i=1:20
   if  hcv(i) < 0.5
       yhcv(i) = 0; %prediccion para el test
   end
end

%% calculamos el error

errorescv=yhcv~=ycv;

errorestr=yhtr~=ytr;

Tetr=sum(errorestr)/80  %error del train

Tecv=sum(errorescv)/20  %error del test

plotDecisionBoundary(theta, Xcv, yhcv);
xlabel('Exam 1 score');
ylabel('Exam 2 score');

%% Probabilidad de ser admitido con un 45 y el resto de valores

primEx = 45*ones(101,1);
segEx = transpose(0:100);
Xpreg = [ones(101,1) primEx segEx];

hpreg = 1./(1+exp(-(Xpreg*theta))); %LA hipotesis CON LA MEJOR THETA sobre todos los datos
yhpreg = ones(size(hpreg));

for i=1:101
   if hpreg(i) < 0.5
       yhpreg(i) = 0;
   end
end

plotDecisionBoundary(theta, Xpreg, yhpreg);
xlabel('Exam 1 score');
ylabel('Exam 2 score');

%% segundo apartado %%

clear ; close all;

%% Load and Plot Data
%  The first two columns contains the X values and the third column
%  contains the label (y).

data = load('mchip_data.txt');
X = data(:, [1, 2]); 
y = data(:, 3);
N = length(y);
p = randperm(N); %reordena aleatoriamente los datos
X = X(p,:);
y = y(p);

plotData(X, y);
xlabel('Microchip Test 1')
ylabel('Microchip Test 2')
legend('y = 1', 'y = 0')

[ Xcv, ycv, Xtr, ytr] = particion(1, 5, X, y);
%% cambio expansion
X = mapFeature(X(:,1), X(:,2));
Xtr = mapFeature(Xtr(:,1), Xtr(:,2));
Xcv = mapFeature(Xcv(:,1), Xcv(:,2));
%% calculo del mejor lambda

[best_size, listError_train, listError_cv] = KfoldCrossValidation(20,X,y);

%% uso el mejor lambda para calcular la hipotesis
theta = transpose(1:28);
options.display = 'final'; %otros: 'iter' , 'none‘
options.method = 'newton'; %por defecto: 'lbfgs'
theta = minFunc(@CosteLogReg, theta, options, Xtr, ytr, best_size);            %minfunc
        
htr = 1./(1+exp(-(Xtr*theta))); %LA hipotesis CON LA MEJOR THETA sobre todos los datos
yhtr = ones(size(htr));

hcv = 1./(1+exp(-(Xcv*theta)));
yhcv = ones(size(hcv));
        
for i=1:size(htr,1)
   if htr(i) < 0.5
     yhtr(i) = 0; %predicción para el train
   end
end

for i=1:size(hcv,1)
    if  hcv(i) < 0.5
       yhcv(i) = 0; %prediccion para el test
    end
end

%% calculamos el error

errorescv=yhcv~=ycv;

errorestr=yhtr~=ytr;

Tetr=sum(errorestr)/size(yhtr,1)  %error del train

Tecv=sum(errorescv)/size(yhcv,1) %error del test

plotDecisionBoundary(theta, Xcv, yhcv);
title(sprintf('lambda = %g', best_size))
xlabel('Microchip Test 1')
ylabel('Microchip Test 2')
legend('y = 1', 'y = 0', 'Decision boundary')

%% ahora lambda = 0
theta = transpose(1:28);
theta = minFunc(@CosteLogReg, theta, options, Xtr, ytr, 0);            %minfunc
        
htr = 1./(1+exp(-(Xtr*theta))); %LA hipotesis CON LA MEJOR THETA sobre todos los datos
yhtr0 = ones(size(htr));

hcv = 1./(1+exp(-(Xcv*theta)));
yhcv0 = ones(size(hcv));
        
for i=1:size(htr,1)
   if htr(i) < 0.5
     yhtr0(i) = 0; %predicción para el train
   end
end

for i=1:size(hcv,1)
    if  hcv(i) < 0.5
       yhcv0(i) = 0; %prediccion para el test
    end
end
%% calculamos el error

errorescv=yhcv0~=ycv;

errorestr=yhtr0~=ytr;

Tetr0=sum(errorestr)/size(yhtr0,1)  %error del train

Tecv0=sum(errorescv)/size(yhcv0,1) %error del test

plotDecisionBoundary(theta, Xcv, yhcv);
title(sprintf('lambda = %g', 0))
xlabel('Microchip Test 1')
ylabel('Microchip Test 2')
legend('y = 1', 'y = 0', 'Decision boundary')

%% tercer apartado calculo de la matriz de confusión

matrizDeConfusion = zeros(3,3);
%encontramos los aciertos y los fallos
TP = find(yhcv==ycv);
fallos = find(yhcv ~= ycv);

for i=1:size(TP)
%sabiendo que son aciertos, sumamos en esa diagonal +1   
    matrizDeConfusion(yhcv(TP(i,1))+1,yhcv(TP(i,1))+1) = matrizDeConfusion(yhcv(TP(i,1))+1,yhcv(TP(i,1))+1) + 1;
end
for i=1:size(fallos)
%sabiendo que son fallos, sumamos el error, en la fila correcta, la columna
%decidida
    matrizDeConfusion(y(fallos(i,1))+1,yhcv(fallos(i,1))+1) = matrizDeConfusion(y(fallos(i,1))+1,yhcv(fallos(i,1))+1) + 1;
end
suma = sum(matrizDeConfusion);
for i=1:2
   matrizDeConfusion(i, 3) = sum(matrizDeConfusion(i,:));
   matrizDeConfusion(3,i)=suma(i);
end
suma = sum(matrizDeConfusion(:,3));
matrizDeConfusion(3,3)=suma;

matrizDeConfusion
%% recall,f1,precision
recall=[];
precision=[];
for i=1:2
    recall = [recall,matrizDeConfusion(i,i)/matrizDeConfusion(i,3)];
    precision = [precision,matrizDeConfusion(i,i)/matrizDeConfusion(3,i)];
    recall = mean(recall);
    precision = mean(precision);
end
recall
precision
f1=(2*(recall*precision))/(recall+precision);
