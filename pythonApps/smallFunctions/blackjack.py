import random
from abc import ABC, abstractmethod

#creates an idividual id for each card based on suit and rank
class Card:
    """
    Attribute
    ----------
    suit : str
        Either heart, diamon, club or spade
    rank : str
        This is the string value of a card or in the case of a face card jack... king etc.
    value: int
        the numeric value of each rank

    """
    def __init__(self, suit, rank):
        self.suit = suit
        self.rank = rank
        self.value = 0

    def __repr__(self):
        """ __repr__ : str : returns the name of every card as s tring of rank and suit """
        return "{rank} of {suit}".format(rank = self.rank, suit = self.suit)


class Deck:
    """
    Attribute
    ----------
    cards : list
        Creates a list with all 52 possible cards

    """
    def __init__(self):
        self.cards = [Card(suit, rank) for suit in ["Spades", "Clubs", "Hearts", "Diamonds"] for rank in ["Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"]]
    
    def deal_card(self):
        """ deal_card : function : takes one card from the list of deck and returns """
        return self.cards.pop(0)
            
    def shuffle(self):
        """shuffle : function : the deck is created everytime in the same order so this randomizes the order"""
        random.shuffle(self.cards)


class Player(ABC):
    """
    Attribute
    ----------
    cards : list
        Creates a list with all 52 possible cards
    score : int
        Tracks the numeric score of each hand for each player

    """

    def __init__(self):
        self.cards = []
        self.score = 0
    
    def __len__(self):
        """len : function : returns the length of the list of cards for any players hand """
        return len(self.cards)
    
    def __iter__(self):
        """iter : function : returns the iterator for the list of cards in any players hand"""
        return iter(self.cards)

    def add_card(self, card):
        """add_card : function : appends a card to the list in the players hand"""
        self.cards.append(card)
    
    def get_score(self):
        """ get_score: 
                function : that calculates the score in a player hands
                returns : int for that score """
        self.score_in_hand()
        return self.score
    
    def bust(self):
        """ bust: function : that determines if the player has more than allowed score of 21
            returns : a bool
            """
        if self.score_in_hand() > 21:
            return True 
        else:
            False

    def score_in_hand(self):
        """ score_in_hand : function : calculates the score of any hand by converting ranks to a numeric values
            returns : int
        """
        self.score = 0
        self.ace_for_blackjack = False
        for card in self.cards:
            if card.rank.isnumeric():
                self.score += int(card.rank)
            else:
                #Aces are worth 11 and required to get black jack
                if card.rank == "Ace":
                    self.ace_for_black_jack = True
                    self.score += 11
                #every other card that is numeric ie Jack through King is worth 10 points
                else:
                    self.score += 10
        return self.score
    
    @abstractmethod
    def print_hand(self):
        """" print_hand : defined as an abstract method and is expanded on in sub classes of Player """
        pass

    @abstractmethod
    def will_hit(self):
        """" will_hit : defined as an abstract method and is expanded on in sub classes of Player """
        pass


class HumanPlayer(Player):
    """
    Attribute: no new attributes all inherited from Player class
    ----------
    """
    def __init__(self):
        super().__init__()
    
    
    def will_hit(self):
        """will_hit : function that allows the player to play black jack by deciding to add 
        cards to their hand (hit) or stay with their current hand (stand)
                parameters: self.score > 21 : if the player has more than 21 they have lost
                and there is no need to ask for the player input
                        choice: must be hit or stand"""
        self.get_score()
        if self.score <= 21: 
            choice = input("Do you want to [hit] or [stand]?").lower()
            while choice not in ["hit","stand"]:
                choice = input("Please enter [hit] or [stand]?").lower()
            if choice == "hit":
                return "hit"
            elif choice == "stand":
                return "stand"
        else:
            return"bust"

    def print_hand(self):
        """print hand: abstract method that the Human Player wants to see all of their cards
        so it will print all of the players cards.  """
        print("Your hand: ")
        for card in self.cards:
            print(card)
        print("Your Score:", self.score_in_hand())    


class ComputerPlayer(Player):
    """
    Attribute: no new attributes all inherited from Player class
    ----------
    """
    def __init__(self):
        super().__init__()
    
    def print_hand(self, player_turn = False):
        """print hand: abstract method that the ComputerPlayer has that hides their first card
            unless it is their turn in which case it prints all of their cards.  """
        if player_turn == False:
            print("computer card 1: hidden")
            print("computer card 2: ",self.cards[0])
        elif player_turn == True:
            for i in self.cards:
                print("computer: ", i)

    def will_hit(self):
        """will_hit : function that allows the computer to play black jack by deciding to add 
        cards to their hand (hit) or stay with their current hand (stand)
                parameters: self.score > 21 : if the computer has more than 21 they have lost
                and there is no need to ask for the player input
                        choice: must be hit or stand"""
        self.get_score()
        if self.score < 19:
            print("computer hits")
            self.print_hand( player_turn = True)
            return "hit"
        elif self.score >= 19:
            print("computer stands")
            self.print_hand( player_turn = True)
            return "stand"


class DealerPlayer(Player):
    """
    Attribute: no new attributes all inherited from Player class
    ----------
    """
    def __init__(self):
        super().__init__()

    
    def print_hand(self, player_turn = False):
        """print hand: abstract method that the DealerPlayer has that hides their first card
            unless it is their turn in which case it prints all of their cards.  """
        if player_turn == False:
            print("dealer card 1: hidden")
            print("dealer card 2: ",self.cards[0])
        elif player_turn == True:
            for i in self.cards:
                print("dealer: ", i)

            
    def will_hit(self):
        """will_hit : function that allows the Dealer to play black jack by deciding to add 
        cards to their hand (hit) or stay with their current hand (stand)
                parameters: self.score > 21 : if the Dealer has more than 21 they have lost
                and there is no need to ask for the player input
                        choice: must be hit or stand"""
        self.get_score()
        if self.score < 17:
            print("dealer hits")
            self.print_hand(player_turn = True)
            return "hit"
        elif self.score >= 17:
            print("dealer stands")
            self.print_hand(player_turn = True)
            return "stand"


class Game:
    """
    This class serves as this inner workings for the blackjack game and has no organic attributes
    instead it assumes the attributes of the other classes that it defines in itself.
    Attributes:
    -----------
    deck
    player_hand :
    dealer_hand : 
    computer_hand :
    """
    def __init__(self):
        pass
    
    def __len__(self):
        """len : function : returns the length of the list of cards for any players hand """
        return len(self.cards)

    def play(self):
        """play contains the bulk of the process of playing black jack
        Paramaters:
        playing set to True to keep the infinite loop commnand line prompted game
        game_over set to False until you lose and then prompted to continue playing with yes or no"""
        playing = True
        
        #Initiate the deck and initiate the hands for each player
        while playing:
            self.deck = Deck()
            self.deck.shuffle()
            self.player_hand = HumanPlayer()
            self.dealer_hand = DealerPlayer()
            self.computer_hand = ComputerPlayer()
            
            #deal out two cards to each player
            for i in range(2):
                self.player_hand.add_card(self.deck.deal_card())
                self.dealer_hand.add_card(self.deck.deal_card())
                self.computer_hand.add_card(self.deck.deal_card())
            
            #Print the cards that are dealt
            self.player_hand.print_hand()
            self.computer_hand.print_hand()
            self.dealer_hand.print_hand()
            
            #set the game over to False to until they lose or win 
            game_over = False
            
            # #set a variable to track who busts and who has not, assign everyone to False until they bust
            # self.computer_ha = False
            # self.dealer_bust = False
            # self.player_bust = False

            #since cards have just been dealt we will check for blackjack
            while game_over == False:
                player_has_blackjack, dealer_has_blackjack, computer_has_blackjack = self.check_for_blackjack()
                if player_has_blackjack or dealer_has_blackjack or computer_has_blackjack == True:
                    if player_has_blackjack == True:
                        self.player_hand.print_hand()
                        print("player has blackjack and wins")
                    elif computer_has_blackjack == True:
                        self.computer_hand.print_hand(player_turn = True)
                        print("computer has blackjack and wins")
                    elif dealer_has_blackjack == True:
                        self.dealer_hand.print_hand(player_turn = True)
                        print("dealer has blackjack and wins")
                    game_over = True
                    continue
                
                def will_player_hit():
                    """broke up the player hitting into a smaller subfunction
                    this allows blackjack to use recursion instead of an additional while loop """
                    nonlocal game_over
                    choice = self.player_hand.will_hit()
                    if choice in ['hit']:
                        self.player_hand.add_card(self.deck.deal_card())
                        self.player_hand.print_hand()
                        if self.player_hand.bust() == True:
                            print("You lose")
                            game_over == True
                        else:
                            return will_player_hit()
                    elif choice in ["stand"]:
                        self.player_hand.bust = False
                        return print("You stand with a score of: ",self.player_hand.score_in_hand())
                    elif choice in ["bust"]:
                        game_over == True
                        return self.player_hand.bust
                will_player_hit()
                
                print("computers turn")
                def will_computer_hit():
                    """ computer_hit also implementd into a subfuction to allow recursion 
                        instead of the anotehr while looop""" 
                    if self.computer_hand.will_hit() in ['hit']:
                        self.computer_hand.add_card(self.deck.deal_card())
                        if self.computer_hand.bust() == True:
                            print("computer player loses")
                        else:
                            return will_computer_hit()
                if self.computer_hand.will_hit() in ['stand']:
                    return self.computer_hand.score_in_hand() 
                will_computer_hit()


                print("dealers turn")
                def will_dealer_hit():
                    """will_dealer_hit also uses its own subfunction instead of a while loop
                    allowing recursion to work its magic
                    """
                    if self.dealer_hand.will_hit() in ['hit']:
                        self.dealer_hand.add_card(self.deck.deal_card())
                        if self.dealer_hand.bust() == True:
                            print("dealer loses")
                        else:
                            will_dealer_hit()
                if self.dealer_hand.will_hit() in ['stand']:
                    return self.dealer_hand.score_in_hand()
                will_dealer_hit()
                
                def who_wins():
                    """ this functions takes all of the previous scores and determines the winner of every round
                    at the end it ends the game 
                    parameters: bust for each player this is the first thing to check because even if you have the largest score if you busted it doesn't matter
                                score this is compared with all of the participants that did not bust keeping in mind that every player is playing against the dealer
                                your hand only has to be better than the dealer"""

                    nonlocal game_over
                #check to see if the dealer busted
                    if not self.dealer_hand.bust:
                        if self.computer_hand.bust and self.player_hand.bust == True:
                            print("dealer wins! Other players bust")
                    #dealer wins all ties in blackjack
                    #if the dealer score is higher than either of ther players the dealer wins
                        if self.dealer_hand.score_in_hand() >= self.computer_hand.score_in_hand() or self.player_hand.score_in_hand():
                            print("dealer wins  with a score of: ", self.dealer_hand.score_in_hand())
                    #if the computer did not bust
                        if self.computer_hand.bust == False:
                        #if both the computer and the player did not bust and both of their scores are higher than the dealer then they win
                            if self.dealer_hand.score_in_hand() < self.computer_hand.score_in_hand() and self.player_hand.score_in_hand():
                                print("player and computer win with scores of: ",self.player_hand.score_in_hand() ,self.computer_hand.score_in_hand() ,"respectively")
                        #if the computer only beats the dealer
                            elif self.dealer_hand.score_in_hand() < self.computer_hand.score_in_hand():
                                print("computer wins with scores of: ",self.computer_hand.score_in_hand())
                        #if the player only beats dealer
                            elif self.dealer_hand.score_in_hand() < self.player_hand.score_in_hand():
                                print("player wins: ",self.player_hand.score_in_hand())
                        #if the computers busts then we only compare player and dealer
                        elif self.computer_hand.bust == True:
                            #if the dealer has an equal or better hand he wins
                            if self.dealer_hand.score_in_hand() >= self.player_hand.score_in_hand():
                                print("dealer wins  with a score of: ", self.dealer_hand.score_in_hand())
                            elif self.dealer_hand.score_in_hand() < self.player_hand.score_in_hand():
                                print("player wins: ",self.player_hand.score_in_hand())
                    

                #You only play against the dealer so if the dealer busted then everybody wins
                    elif self.dealer_hand.bust:
                        if self.computer_hand.bust and self.player_hand.bust == False:
                            print("The computer and player win!!!!")
                            print("Computer hand:", self.computer_hand.score_in_hand())
                            print("Player hand:", self.player_hand.score_in_hand())
                        elif self.computer_hand.bust == False:
                            print("The computer wins!")
                            print("Computer hand:", self.computer_hand.score_in_hand())
                    elif self.dealer_hand.score_in_hand() and self.computer_hand.score_in_hand() and self.player_hand.score_in_hand() == True:
                        print("Everyone busts and looses!!!")
                    game_over = True
                who_wins()
                
            again = input("Would you like to keep losing money [Y/N] ")
            while again.lower() not in ["y", "n"]:
                again = input("Try again, enter Y or N ")
            if again.lower() == "n":
                print("Please come back and lose more money next time")
                playing = False
            else:
                game_over = False



    def check_for_blackjack(self):
        """ check_for_blackjack determines if anyone got lucky with their first two cards and immediately ends the game
        parameters: are set for each player and are initially set to False
        if True they return and print in play who has won"""
        player = False
        dealer = False
        computer = False
        if self.player_hand.score_in_hand() == 21:
            player = True
        if self.dealer_hand.score_in_hand() == 21:
            dealer = True
        if self.computer_hand.score_in_hand() == 21:
            computer = True

        return(player, dealer, computer)



if __name__ == "__main__":
    g = Game()
    g.play()