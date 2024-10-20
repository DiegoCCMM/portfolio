function [best_size, listError_train, listError_cv] = KfoldCrossValidation(k,X,y)
%KFOLDCROSSVALIDATION Summary of this function goes here
%   Detailed explanation goes here

best_size = 0; best_errV = inf;

theta = transpose(1:28);

listError_train = []; 
listError_cv = [];

for lambda = logspace(-10,0) 
    err_T = 0; err_V = 0;
    
    for fold = 1:k 
        %{separar N/k ejemplos para validacion}
        [ Xcv, ycv, Xtr, ytr] = particion(fold, k, X, y);
        
        
        options.display = 'final'; %otros: 'iter' , 'none‘
        options.method = 'newton'; %por defecto: 'lbfgs'
        theta = minFunc(@CosteLogReg, theta, options, Xtr, ytr, lambda);            %minfunc
        
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
        
        numErrorcv = length(find(ycv ~= yhcv));                           %los errores
        numErrorTr = length(find(ytr ~= yhtr));       
            
        err_V = err_V + numErrorcv/length(ycv);                             %acumulacion de tasas
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

