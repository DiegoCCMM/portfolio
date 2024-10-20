function Z = updateClusters(D,mu,k)
% D(m,n), m datapoints, n dimensions
% mu(K,n) final centroids
m = size(D,1);
Z = ones(1,m);

for i=1:m
    min = ((norm(D(i,:) - mu(1,:))).^2);
    for j=2:k
        
        candidato=((norm(D(i,:) - mu(j,:))).^2);
        
        if(prod(candidato<min))
            Z(i) = j;
            min=candidato;
        end
        
    end
    
end
% c(m) assignment of each datapoint to a class
end
