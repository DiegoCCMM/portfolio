function yhat = clasificacionBayesiana(modelo, nc, X, clases)
% Con los modelos entrenados, predice la clase para cada muestra X
% modelo{i}.N     : Numero de muestras de la clase i
% modelo{i}.mu    : Media de la clase i
% modelo{i}.Sigma : Covarianza de la clase i

probabilidadPorClases = [];

[N,m] = size(X);

yhat=zeros(N,1);

for i=1:nc
    cond=ismember(i,clases);
    if(cond(1))
        probabilidadPorClases=[probabilidadPorClases gaussLog(modelo{i}.mu, modelo{i}.Sigma, X)];
    end
end
%ahora en probabilidadPorClases tenemos la probabilidad de que un dato, de
%una fila pertenezca a una clase, que están en cada columna (habiendo 10
%columnas)
for i=1:N
    %encontramos el indice en donde está el valor más, es decir, a que
    %clase tiene más probabilidad de pertencer este dato
    [max_num, max_idx] = max(probabilidadPorClases(i,:));
    [X, Y]=ind2sub(size(X),max_idx);
    
    yhat(i,1) = Y;
    %una vez tenemos a que clase pertenece la almacenamos de forma que
    %tenemos a que clase corresponde el dato de la fila.
end

end