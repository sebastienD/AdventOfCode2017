import unittest
from main import analyzeLine, initReg

#  * pour créer un test unitaire, il faut créer une classe qui hérite deunittest.TestCase. 
#  * Les méthodes de test ont un nom commençant par test
#  * python3 -m unittest: permet la découverte automatique des tests dans le répertoire courant.
class MainTest(unittest.TestCase):

    """Test case utilisé pour tester les fonctions du module 'random'."""

    #def setUp(self):
    #    """Initialisation des tests."""

    def test_zero(self):
        line = "b inc 5 if a > 1"
        registers = analyzeLine(line)
        self.assertEqual(registers["a"], 0)
        self.assertEqual(registers["b"], 0)

    def test_inc(self):
        line = "a inc 1 if b < 5"
        registers = analyzeLine(line)
        self.assertEqual(registers["a"], 1)
        self.assertEqual(registers["b"], 0)

# Ceci lance le test si on exécute le script
# directement.
if __name__ == '__main__':
    unittest.main()