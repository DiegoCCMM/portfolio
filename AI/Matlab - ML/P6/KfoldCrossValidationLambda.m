function [best_size, listError_train, listError_cv] = KfoldCrossValidationLambda(k,X,y,bayesIngenuo,clases)
%KFOLDCROSSVALIDATION Summary of this function goes here
%   Detailed explanation goes here

best_size = 0; best_errV = inf;

listError_train = []; 
listError_cv = [];

for lambda = logspace(-10,0) 
    err_T = 0; err_V = 0;
    
    for fold = 1:k 
        %{separar N/k ejemplos para validacion}
        [ Xcv, ycv, Xtr, ytr] = particion(fold, k, X, y);
        
        
        modelo = entrenarGaussianas( Xtr, ytr, 10, bayesIngenuo, lambda );
        
        yhatcv = clasificacionBayesiana(modelo, 10, Xcv,clases);
        yhatTr = clasificacionBayesiana(modelo, 10, Xtr,clases);
        
        numErrorcv = length(find(ycv ~= yhatcv));
        numErrorTr = length(find(ytr ~= yhatTr));
        
        err_V = err_V + numErrorcv/length(ycv);
        err_T = err_T + numErrorTr/length(ytr);

    end
    err_T = err_T/k; err_V = err_V/k;
    %{calcular el error medio de las k veces}
    if err_V < best_errV
        best_size = lambda;
        best_errV = err_V;
        %{guardar el mejor valor de los hyper-parametros}
    end
    listError_train = [listError_train, err_T];
    listError_cv = [listError_cv, err_V];
end
end

