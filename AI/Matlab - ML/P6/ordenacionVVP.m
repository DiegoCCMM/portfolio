function [Ures,lambdares] = ordenacionVVP(U,lambda)
%ORDENACIONVVP ordena los valores propios de mayor a menor y los vectores
%propios en ese mismo orden
    [n,m]=size(lambda);
    lambdares=ones(n,m);
    Ures=ones(n,m);
    listaMaximos=[];
    ordenVectores = [];
    lambdaaux=max(lambda);
    for i=1:m
       [maximo,j]=max(lambdaaux);
       listaMaximos=[listaMaximos,maximo];
       
       ordenVectores = [ordenVectores,j];
       lambdaaux(j)=[];
    end
    
    lambdares=diag(listaMaximos);
    
    for i=1:m
        Ures(:,i)=U(:,ordenVectores(i));
    end
end

