function [best_size, listError_train, listError_cv] = KfoldCrossValidation(k,X,y)
%KFOLDCROSSVALIDATION Summary of this function goes here
%   Detailed explanation goes here

best_size = 0; best_errV = inf;

listError_train = []; 
listError_cv = [];
theta0 = zeros(size(X,2),1);

options.display = 'final'; %otros: 'iter' , 'none‘
options.method = 'pnewton0'; %por defecto: 'lbfgs'
options.useMex = 1;

for lambda = logspace(-10,0) 
    err_T = 0; err_V = 0;
    
    for fold = 1:k
        
        THETA = [];
        %{separar N/k ejemplos para validacion}
        [ Xcv, ycv, Xtr, ytr] = particion(fold, k, X, y);
        for j = 1:10
            ytrj=ytr==j;
            theta = minFunc(@CosteLogReg, theta0, options, Xtr, ytrj, lambda);            %minfunc
            THETA = [THETA, theta];
        end
        htr = 1./(1+exp(-(Xtr*THETA))); %LA hipotesis CON LA MEJOR THETA sobre todos los datos
        yhtr = zeros(size(htr,1),1);

        hcv = 1./(1+exp(-(Xcv*THETA)));
        yhcv = zeros(size(hcv,1),1);

        for i=1:size(htr,1)
            [val, loc] = max(htr(i,:));
            yhtr(i) = loc;
        end

        for i=1:size(hcv,1)
            [val, loc] = max(hcv(i,:));
            yhcv(i) = loc;
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

