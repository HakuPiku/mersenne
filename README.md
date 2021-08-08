# mersenne


Mersenne Twister(MT19937-32) implementation based on  https://en.wikipedia.org/wiki/Mersenne_Twister


Shouldn't be used anywhere. Never Roll Your Own Crypto. If you want a random number generator use Go's native math/rand library. Unless you want it to be unpredictable, in that case you need a cryptographically secure pseudo-random number generator(CSPRNG), so use crypto/rand. 

