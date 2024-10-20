function [k] = elegirK(lambda, variabilidad)
%ELEGIRK Summary of this function goes here
%   Detailed explanation goes here
[n,m]=size(lambda);
k=0;
valoresPropios=max(lambda);

    for i=1:m
        valVar=((sum(valoresPropios(1:i)))/sum(valoresPropios));
        if((valVar*100)>=variabilidad && k==0)
                k=i;
        end
    end
end

