module MariosMarvellousLasagna exposing (remainingTimeInMinutes)

remainingTimeInMinutes layers timeAlreadySpent =
    let
        preparationTime = layers * 2
        timeInOven = 40
    in
    
    preparationTime + timeInOven - timeAlreadySpent