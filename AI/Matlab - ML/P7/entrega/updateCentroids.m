function munew = updateCentroids(D,c)
% D((m,n), m datapoints, n dimensions
% c(m) assignment of each datapoint to a class
m = size(D,1);
munew=zeros(16,3);
for k=1:16
    
    numPunto = 0;
    for i=1:m
        
        if(c(i) == k)
            munew(k,:)=munew(k,:)+D(i,:);
            numPunto = numPunto+1;
        end
        
    end
    munew(k,:)=munew(k,:)./numPunto;
    
end
% munew(K,n) new centroids
end
