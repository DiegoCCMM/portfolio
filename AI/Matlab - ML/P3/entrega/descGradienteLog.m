function [ListJ,it,theta] = descGradienteLog(theta, X, y)
% calculo del descenso del gradiente logistico

%algoritmo de descenso de gradiente con el coste logistico

ListJ = []; it = 0;
%J con la theta
[J0,g]=CosteLogistico(theta,X,y);

ListJ = [ListJ, J0]; it = it +1;
%J1 con la theta
theta = theta - 0.001*g;
[J1,g]=CosteLogistico(theta,X,y);

ListJ = [ListJ, J1];it = it +1;

while( J0 - J1 > 0.00001)
    %J1 será el inicial
    J0 = J1;
    %con el gradiente de J0 que fue J1, obtendremos J2 = J1
    theta = theta - 0.001*g;
    [J1,g]=CosteLogistico(theta,X,y);
    
    ListJ = [ListJ, J1];
    it = it + 1;
end

end

