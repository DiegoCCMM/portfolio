function [mu, c] = kmeans(D,mu0,K)

% D(m,n), m datapoints, n dimensions
% mu0(K,n) K initial centroids

m = size(D,1);
n = size(D,2);
bool = false;
    while bool == false
         c=updateClusters(D, mu0,K);
         
         mu=updateCentroids(D, c);
        
         bool=isequal(mu,mu0);
         mu0=mu;
    end
end
% mu(K,n) final centroids
% c(m) assignment of each datapoint to a class
