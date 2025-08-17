module CollatzConjecture exposing (collatz)


collatz : Int -> Result String Int
collatz start =
    if start < 1 then
        Err "Only positive integers are allowed"

    else
        Ok <| iterate_collatz 0 start

iterate_collatz steps x =
    let
        next =
            iterate_collatz (steps + 1)
    in

    if x == 1 then
        steps

    else if modBy 2 x == 0 then
        next (x // 2)

    else
        next (3 * x + 1)