function [ X, mu ] = estandarizar( X )
% xij - muj
     mu = mean ( X );
     
     for i=1:size(X,1)
         for j=1:size(X,2)
             
             X(i,j)=X(i,j)-mu(j);
             
         end
     end


end