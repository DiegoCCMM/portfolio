function [theta] = regresionLogReg(X,y, lambda)
%UNTITLED codigo necesario de la practica 3 para hacerlo multiclase

%% expansion
X = mapFeature(X(:,1), X(:,2));
%% uso el mejor lambda para calcular la hipotesis
theta = transpose(1:28);
options.display = 'final'; %otros: 'iter' , 'none‘
options.method = 'newton'; %por defecto: 'lbfgs'
theta = minFunc(@CosteLogReg, theta, options, X, y, lambda);            %minfunc


end

