from setuptools import setup, find_packages

setup (
    author='geremachek',
    author_email='gmk@airmail.cc',
    name='flip',
    long_description='Flip through files',
    packages=find_packages(),
    entry_points={'console_scripts': ['flip = flip.__main__:main']},
    package_dir={'flip': 'flip'},
)
