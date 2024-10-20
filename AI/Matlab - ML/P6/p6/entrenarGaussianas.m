function modelo = entrenarGaussianas( Xtr, ytr, nc, NaiveBayes, landa )
% Entrena una Gaussana para cada clase y devuelve:
% modelo{i}.N     : Numero de muestras de la clase i
% modelo{i}.mu    : Media de la clase i
% modelo{i}.Sigma : Covarianza de la clase i
% Si NaiveBayes = 1, las matrices de Covarianza ser�n diagonales  (bayes
% ingenuo)
% Se regularizar�n las covarianzas mediante: Sigma = Sigma + landa*eye(D)

[m,D] = size(Xtr);

for i=1:nc
    %calculo del número de elementos de una clase
    elementos=find(ytr==i);
    modelo{i}.N=numel(elementos);
    %calculo de la media de los valores de los pixeles de una clase
    %(clases,muPixeles)
    modelo{i}.mu=mean(Xtr(elementos,:));
    %calculo de la covarianza de una clase
    modelo{i}.Sigma=cov(Xtr(elementos,:));
    %regularización para evitar sobreajuste
    modelo{i}.Sigma=modelo{i}.Sigma + landa*eye(D);
    
    if(NaiveBayes==1)
    %versión BayesIngenuo con la matriz diagonal (el resto de valores
    %fuera de la diagonal 0)
        modelo{i}.Sigma=diag(diag(modelo{i}.Sigma));
    end
    
end


