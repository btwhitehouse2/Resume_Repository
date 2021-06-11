import re
from itertools import chain, permutations
import problem1 as p1

def create_scrabble_rack():
    """confirms that the input is only letters, has 
    more than 0 and less than 8. Removes all spaces and commas.
     """
    while True:
        scrabble_rack = input("scrabble rack: ")
        scrabble_rack = scrabble_rack.replace(',','')
        scrabble_rack = scrabble_rack.replace(' ','')
        if len(scrabble_rack) < 8:
            if scrabble_rack.isalpha():
                break
            else:
                print("Please enter characters a-z only")
        else:
            print("Please only enter seven tiles")
    return(scrabble_rack)


def list_of_outcomes(rack):
    """takes input of scrabble rack and creates all possible permutations """    
    temp_list_of_strings = []
    list_of_permutations = []
    lo, hi = 2, 7

    for perm in chain.from_iterable(permutations(rack, i) for i in range(lo, hi + 1)):
        temp_list_of_strings = (''.join(perm))
        list_of_permutations.append(temp_list_of_strings)
    return(list_of_permutations)

def import_scrabble_list():
    """ takes scrabble words from list of scrabble words and inputs to list """
    scrable_file = open("scrabble_list.txt", "r")

    list_of_possible_scrabble_words = []
    for line in scrable_file:
        word_stripped_line = line.strip()
        word_stripped_line = word_stripped_line.lower()
        list_of_possible_scrabble_words.append(word_stripped_line)

    return(list_of_possible_scrabble_words)

def matching_words(list_of_permutations,list_of_possible_scrabble_words):
    """takes both the list of permutations and compares to the list of words and creates a list of words found in both lists"""
    list_of_matching_words=[]
    for i in range(0,len(list_of_possible_scrabble_words)):
        if list_of_possible_scrabble_words[i]  in list_of_permutations:
            list_of_matching_words.append(list_of_possible_scrabble_words[i])
    return(list_of_matching_words)

def organize_words_by_length (scrabble_options):
    """this functions takes the overall list of matching words
    and creates an individual sublist of each length of word from
    2 to 7. It returns these seven lists to main so they can be printed with specific requirments"""
    
    #This part of the code uses regular expression to match terms to each word length sublist    
    scrabble_options = str(scrabble_options)
    #print(scrabble_options)
    two_lwords =  re.findall(r"'\w{2}'\S\s\d+",scrabble_options)
    three_lwords =  re.findall(r"'\w{3}'\S\s\d+",scrabble_options)
    four_lwords =  re.findall(r"'\w{4}'\S\s\d+",scrabble_options)
    five_lwords =  re.findall(r"'\w{5}'\S\s\d+",scrabble_options)
    six_lwords =  re.findall(r"'\w{6}'\S\s\d+",scrabble_options)
    seven_lwords =  re.findall(r"'\w{7}'\S\s\d+",scrabble_options)

    if not two_lwords:
        two_lwords = "no words"
    if not three_lwords:
        three_lwords = "no words"
    if not four_lwords:
        four_lwords = "no words"
    if not five_lwords:
        five_lwords = "no words"
    if not six_lwords:
        six_lwords = "no words"
    if not seven_lwords:
        seven_lwords = "no words"

    return(two_lwords,three_lwords,four_lwords,five_lwords,six_lwords,seven_lwords)

def main():
    """in this part of the problem imports parts from problem1 in order to get to the final list with scores """
    scrabble_rack = create_scrabble_rack()
    list_of_permutations = list_of_outcomes(scrabble_rack)
    list_of_possible_scrabble_words = import_scrabble_list()
#    print(list_of_possible_scrabble_words)
    list_of_matching_words = matching_words(list_of_permutations,list_of_possible_scrabble_words)

    #this piece of codes takes the matching list and applies to problem1 to get the score for those applicable words
    score_dict = p1.scrabble_score(list_of_matching_words)
    #this pieces of code creates a dictionary from the list based off of the scores and the words
    #this sorts the dictionary based off of the word lengths
    scrabble_options = p1.dictionary_sort(score_dict)
    two_lwords,three_lwords,four_lwords,five_lwords,six_lwords,seven_lwords = organize_words_by_length (scrabble_options)
    #This final print statement prints all of the sublists with the 10 highest scoring terms
    #Includes spacing and character sperators so that the user does not become confused
    print(r'2 Letter words','\n','-'*60,'\n', *two_lwords[:10],'\n'*5,
    '3 letter_words','\n','-'*60,'\n', *three_lwords[:10],'\n'*5, 
    '4 letter words','\n','-'*60,'\n',*four_lwords[:10],'\n'*5,
    '5 letter words','\n','-'*60,'\n',*five_lwords[:10],'\n'*5,
    '6 letter words','\n','-'*60,'\n',*six_lwords[:10],'\n'*5,
    '7 letter words','\n','-'*60,'\n', *seven_lwords[:10])    


if __name__ == "__main__":
    main()









""" 
Below I coded different ways to compare two different lists
"""

"""
def matching_words(list_of_permutations,list_of_possible_scrabble_words):
    
    list_of_matching_words=[]
    for i in range(0,len(list_of_possible_scrabble_words)):
        for j in range(0,len(list_of_permutations)):
            if list_of_possible_scrabble_words[i] == list_of_permutations[j]:
                list_of_matching_words.append(list_of_possible_scrabble_words[i])
    return(list_of_matching_words)
"""
"""
def matching_words(list_of_permutations,list_of_possible_scrabble_words):
    scrabble_set = set(list_of_possible_scrabble_words)
    permutation_set = set(list_of_permutations)

    if (scrabble_set & permutation_set):
        print(scrabble_set & permutation_set)
    else:
        print("No common letters")
"""

"""def matching_words(list_of_permutations,list_of_possible_scrabble_words):
    list_of_matching_words=[]
    for word in list_of_possible_scrabble_words:
        if word in list_of_permutations:
            list_of_matching_words.append(word)
    return(list_of_matching_words)
"""
