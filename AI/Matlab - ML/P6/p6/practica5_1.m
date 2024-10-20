function [] = practica5_1(Xgorro,ySelected,Xtest,ytest,clases)
%% encontramos los modelos gaussianos
[best_lambda, listError_train, listError_cv] = KfoldCrossValidationLambda(5,Xgorro,ySelected,1,clases);

%% Re-entrena con todos los datos
modelo = entrenarGaussianas( Xgorro, ySelected, 10, 1, best_lambda );
%% Predicción y clasificación
yCla = clasificacionBayesiana1(modelo, 10, Xtest,clases);
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
    matrizDeConfusion(ySelected(fallos(i,1)),yCla(fallos(i,1))) = matrizDeConfusion(ySelected(fallos(i,1)),yCla(fallos(i,1))) + 1;
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
    cond=ismember(i,clases);
    if(cond(1))
        recall = [recall,matrizDeConfusion(i,i)/matrizDeConfusion(i,11)];
        precision = [precision,matrizDeConfusion(i,i)/matrizDeConfusion(11,i)];
        recall = mean(recall);
        precision = mean(precision);
    end
end
recall
precision
f1=(2*(recall*precision))/(recall+precision);
f1
verConfusiones(Xtest,ytest,yCla);
end

