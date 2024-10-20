% load images 
% images size is 20x20. 
clear
close all

load('MNISTdata2.mat'); 

nrows=20;
ncols=20;

nimages = size(X,1);

% Show the images
for I=1:40:nimages, 
    imshow(reshape(X(I,:),nrows,ncols))
    pause(0.1)
end


%% Perform PCA over all numbers

% z should contain the projections over the first two PC
% now is just a random matrix
z=rand(size(X,1),2);

% Muestra las dos componentes principales
figure(100)
clf, hold on
plotwithcolor(z(:,1:2), y);

%% Use classifier from previous labs on the projected space







