# Game App Requirements

# Use Case

## user :

### register
user can register on the game

### login
user can log in into the game

### Game
- Each game has 4 number of questions
- Each game has 2 players (For now)
- The difficulty of each game is one of "Hard , Medium , Easy"
- The Winner of each Game is the player who has more correct answers

### Player
- The player is defined for each game
- The player score is the count of correct answers

# entity
## user
- ID 
- Name
- Phone number
- Avatar
## Game
- ID
- Category
- Question list
- Difficulty
## Question
- ID
- Category
- Text
## Player
- ID
- User
- Game
- Answer List
- Score
- Is Winner
## Answer
- ID
- Text
- Question
- Difficulty
- Is Correct 