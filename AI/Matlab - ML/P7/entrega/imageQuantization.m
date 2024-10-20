clear all;

figure(1)
im = imread('smallparrot.jpg');
imshow(im)

%% datos
D = double(reshape(im,size(im,1)*size(im,2),3));

%% dimensiones
m = size(D,1);
n = size(D,2);

%% Kmeans 
K = 16;

%% Inicializar centroides
for i=1:K
    mu0(i,:) = D(i*11500,:);
end
%% bucle kmeans
[mu, c] = kmeans(D, mu0, K);
%% reconstruir imagen
qIM=zeros(length(c),3);
for h=1:K,
    ind=find(c==h);
    qIM(ind,:)=repmat(mu(h,:),length(ind),1);
end
qIM=reshape(qIM,size(im,1),size(im,2),size(im,3));
figure(2)
imshow(uint8(qIM));
%% mismo con mi foto
clear all;

figure(1)
im = imread('pilar.jpg');
imshow(im)

%% datos
D = double(reshape(im,size(im,1)*size(im,2),3));

%% dimensiones
m = size(D,1);
n = size(D,2);

%% Kmeans 
K = 16;

%% Inicializar centroides
for i=1:K
    mu0(i,:) = D(i*11500,:);
end
%% bucle kmeans
[mu, c] = kmeans(D, mu0, K);
%% reconstruir imagen
qIM=zeros(length(c),3);
for h=1:K,
    ind=find(c==h);
    qIM(ind,:)=repmat(mu(h,:),length(ind),1);
end
qIM=reshape(qIM,size(im,1),size(im,2),size(im,3));
figure(2)
imshow(uint8(qIM));