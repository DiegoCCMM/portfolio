'''
Ejemplo de entrenamiento de un percepton para reconocer MNIST
v1.0 14-Dic-2018
Juan D. Tardos
'''

import keras
from keras.datasets import mnist
from keras.models import Sequential
from keras.layers import Dense, Dropout, Flatten
from keras.layers import Conv2D, MaxPooling2D
from keras import backend as K
from keras.optimizers import RMSprop, Adam, SGD
from keras.callbacks import EarlyStopping
import time
# from sklearn.metrics import confusion_matrix
import numpy as np
# import matplotlib.pyplot as plt
import P5_utils



def printingX(x):
    for elem in x:
        print(elem + '\n')


verbose = True

print('Loading MNIST dataset...')
# Problem dimensions
img_rows, img_cols = 28, 28
num_pixels = img_rows * img_cols
num_classes = 10
# The data, split between train and test sets
(x_train, y_train), (x_test, y_test) = mnist.load_data()
x_train = x_train.reshape(60000, num_pixels)
x_test = x_test.reshape(10000, num_pixels)
x_train = x_train.astype('float32') / 255
x_test = x_test.astype('float32') / 255
print(x_train.shape[0], 'train samples')
print(x_test.shape[0], 'test samples')
# convert class vectors to binary class matrices
y_train = keras.utils.to_categorical(y_train, num_classes)
y_test = keras.utils.to_categorical(y_test, num_classes)
# Random permutation of training data
np.random.seed(0)
p = np.arange(x_train.shape[0])
np.random.shuffle(p)
x_train = x_train[p]
y_train = y_train[p]

# Función para parar cuando ya no mejora el error de validacion
earlystop = EarlyStopping(monitor='val_loss', patience=5,
                          verbose=1, mode='auto')

x = []

for actv3 in ['empty', 'relu']:
    for actv2 in ['empty', 'relu']:
        for actv in ['softmax']:
            for comp in ['categorical_crossentropy', 'mean_squared_error']:
                for optim in ['RMSprop','Adam']:
                    for neuronas in [16]:
                        for neuronas2 in [16]:
                   # Perceptron de un solo nivel
                            model = Sequential()
                        #        kernel_regularizer=regularizers.l2(0.01)))
                # tercera capa
                            if(actv3 != 'empty'):
                                model.add(Dense(units=neuronas, activation=actv3, input_shape=(num_pixels,)))

                # dos capas
                                model.add(Dropout(0.2))

                            if(actv2 != 'empty'):
                                if actv3 != 'empty':
                                    model.add(Dense(units=neuronas2, activation=actv2))
                                else:
                                    model.add(Dense(units=neuronas2, activation=actv2, input_shape=(num_pixels,)))
                # dropout
                                model.add(Dropout(0.2))

                            if actv2 == 'empty' and actv3 == 'empty':
                                model.add(Dense(units=10, activation=actv, input_shape=(num_pixels,)))
                            else:
                                model.add(Dense(units=10, activation=actv))

                            model.compile(loss=comp,
                              optimizer=optim,
                              metrics=['accuracy'])

                            model.summary()

                            print('Training the NN....')
                            t0 = time.clock()
                            history = model.fit(x_train, y_train,
                                batch_size=128,
                                epochs=20,
                                validation_split=0.1,
                                #callbacks=[earlystop],
                                verbose=verbose)

                            train_time = time.clock() - t0
            # print('**', actv, '  ', optim, '  ', comp, '**')
            # print('%s %.3f%s' % ('Training time: ', train_time, 's'))
            # P5_utils.plot_history(history)

            # Evaluar la red
                            train_score = model.evaluate(x_train, y_train, verbose=0)
                            test_score = model.evaluate(x_test, y_test, verbose=0)
            # print('%s %2.2f%s' % ('Accuracy train: ', 100 * train_score[1], '%'))
            # print('%s %2.2f%s' % ('Accuracy test:  ', 100 * test_score[1], '%'))

                            x.append('**' + actv + '  ' + actv2 + '  ' + actv3 + '  ' + optim + '  '
                                 + comp + '  ' + str(neuronas) + str(neuronas2) + '**')
                            x.append('Training time: ' + str(train_time))
                            x.append('Accuracy train: ' + str(100 * train_score[1]))
                            x.append('Accuracy test:  ' + str(100 * test_score[1]))

printingX(x)

x = []

for actv3 in ['relu']:
    for actv2 in ['relu']:
        for actv in ['softmax']:
            for comp in ['categorical_crossentropy', 'mean_squared_error']:
                for optim in ['RMSprop','Adam']:
                    for neuronas in [32, 64, 128]:
                        for neuronas2 in [32, 64, 128]:
                   # Perceptron de un solo nivel
                            model = Sequential()
                        #        kernel_regularizer=regularizers.l2(0.01)))
                # tercera capa
                            if(actv3 != 'empty'):
                                model.add(Dense(units=neuronas, activation=actv3, input_shape=(num_pixels,)))

                # dos capas
                                model.add(Dropout(0.2))

                            if(actv2 != 'empty'):
                                if actv3 != 'empty':
                                    model.add(Dense(units=neuronas2, activation=actv2))
                                else:
                                    model.add(Dense(units=neuronas2, activation=actv2, input_shape=(num_pixels,)))

                # dropout
                                model.add(Dropout(0.2))
                            if actv2 == 'empty' and actv3 == 'empty':
                                model.add(Dense(units=10, activation=actv, input_shape=(num_pixels,)))
                            else:
                                model.add(Dense(units=10, activation=actv))

                            model.compile(loss=comp,
                              optimizer=optim,
                              metrics=['accuracy'])

                            model.summary()

                            print('Training the NN....')
                            t0 = time.clock()
                            history = model.fit(x_train, y_train,
                                batch_size=128,
                                epochs=20,
                                validation_split=0.1,
                                #callbacks=[earlystop],
                                verbose=verbose)

                            train_time = time.clock() - t0
            # print('**', actv, '  ', optim, '  ', comp, '**')
            # print('%s %.3f%s' % ('Training time: ', train_time, 's'))
            # P5_utils.plot_history(history)

            # Evaluar la red
                            train_score = model.evaluate(x_train, y_train, verbose=0)
                            test_score = model.evaluate(x_test, y_test, verbose=0)
            # print('%s %2.2f%s' % ('Accuracy train: ', 100 * train_score[1], '%'))
            # print('%s %2.2f%s' % ('Accuracy test:  ', 100 * test_score[1], '%'))

                            x.append('**' + actv + '  ' + actv2 + '  ' + actv3 + '  ' + optim + '  '
                                 + comp + '  ' + str(neuronas) + str(neuronas2) + '**')
                            x.append('Training time: ' + str(train_time))
                            x.append('Accuracy train: ' + str(100 * train_score[1]))
                            x.append('Accuracy test:  ' + str(100 * test_score[1]))

printingX(x)

batch_size = 128
num_classes = 10
epochs = 12

# input image dimensions
img_rows, img_cols = 28, 28

# the data, split between train and test sets
(x_train, y_train), (x_test, y_test) = mnist.load_data()

if K.image_data_format() == 'channels_first':
    x_train = x_train.reshape(x_train.shape[0], 1, img_rows, img_cols)
    x_test = x_test.reshape(x_test.shape[0], 1, img_rows, img_cols)
    input_shape = (1, img_rows, img_cols)
else:
    x_train = x_train.reshape(x_train.shape[0], img_rows, img_cols, 1)
    x_test = x_test.reshape(x_test.shape[0], img_rows, img_cols, 1)
    input_shape = (img_rows, img_cols, 1)

x_train = x_train.astype('float32')
x_test = x_test.astype('float32')
x_train /= 255
x_test /= 255
print('x_train shape:', x_train.shape)
print(x_train.shape[0], 'train samples')
print(x_test.shape[0], 'test samples')

# convert class vectors to binary class matrices
y_train = keras.utils.to_categorical(y_train, num_classes)
y_test = keras.utils.to_categorical(y_test, num_classes)

model = Sequential()
model.add(Conv2D(32, kernel_size=(3, 3),
                 activation='relu',
                 input_shape=input_shape))
model.add(Conv2D(64, (3, 3), activation='relu'))
model.add(MaxPooling2D(pool_size=(2, 2)))
model.add(Dropout(0.25))
model.add(Flatten())
model.add(Dense(128, activation='relu'))
model.add(Dropout(0.5))
model.add(Dense(num_classes, activation='softmax'))

model.compile(loss=keras.losses.categorical_crossentropy,
              optimizer=keras.optimizers.Adadelta(),
              metrics=['accuracy'])

model.fit(x_train, y_train,
          batch_size=batch_size,
          epochs=epochs,
          verbose=1)
score = model.evaluate(x_test, y_test, verbose=0)
print('Test loss:', score[0])
print('Test accuracy:', score[1])

# for actv in ['softmax', 'sigmoid']:
#     for comp in ['mean_squared_error', 'categorical_crossentropy']:
#         for optim in ['SGD', 'RMSprop', 'Adam']:
#             # Perceptron de un solo nivel
#             model = Sequential()
#
#     #una capa
#             model.add(Dense(10, activation=actv, input_shape=(num_pixels,)))
#
#
#
#             model.compile(loss= comp,
#               optimizer=optim,
#               metrics=['accuracy'])
#
#             model.summary()
#
#             print('Training the NN....')
#             t0 = time.clock()
#             history = model.fit(x_train, y_train,
#                     batch_size=128,
#                     epochs=20,
#                     validation_split=0.1,
#                     callbacks=[earlystop],
#                     verbose=verbose)
#
#             train_time = time.clock() - t0
#             print('**', actv, '  ', optim, '  ', comp, '**')
#             print('%s %.3f%s' %  ('Training time: ', train_time, 's') )
# #P5_utils.plot_history(history)
#
# # Evaluar la red
#             train_score = model.evaluate(x_train, y_train, verbose=0)
#             test_score = model.evaluate(x_test, y_test, verbose=0)
#             print('%s %2.2f%s' % ('Accuracy train: ', 100*train_score[1], '%' ))
#             print('%s %2.2f%s' % ('Accuracy test:  ', 100*test_score[1], '%'))

# y_pred = model.predict(x_test)
# P5_utils.plot_mnist_confusion_matrix(y_test, y_pred)
