from setuptools import setup, find_packages

with open('requirements.txt', 'r') as requirements:
    requires = [line.strip() for line in requirements]

setup(
    name='squareservice',
    version='0.0.1',
    url='https://github.com/brymck/gke-site/src/squareservice',
    author='Bryan McKelvey',
    author_email='bryan.mckelvey@gmail.com',
    description='Square service',
    packages=find_packages(),
    install_requires=requires,
    entry_points={
        'console_scripts': ['squareservice=squareservice:start'],
    },
)
