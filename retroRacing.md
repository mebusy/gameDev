
# Retro Racing Game

- How to represent track 
    - trackWidth := y / (screenH/2), for `0 <= y <= screenH/2`
    - ![](imgs/retro_racing_0.png)

- How to represent car is moving ?
    - use different grass color
    - but the color should not be distributed uniformly. It should be looked like prespectively.
    - ![](imgs/retro_racing_1.png)
    - We can achieve thie line positioning by using some sine function.
    - ![](imgs/retro_racing_2.png)
    - the x-direction represents the perspective.
    - d is the phase, the phase represents the distance the car is moved around the track.
    - in game, we may modify the function to `sin(20*(1-x)³ + 0.1*d)`
    - ![](imgs/retro_racing_3.png)

- How to define a track ?
    - break the track into discrete sections , and label each section with curvature and distance. 
    - ![](imgs/retro_racing_4.png)
    - as the player moves around the track, we look how far they've traveled. Accumulate the distance for all of the sections and work out which section they are in.  and all we need to display on the screen is the tracks curvature. 

- How to represent track curvature
    - calculate which section you are in , and what is the curvature in this section
    - update the current track curvature with  
        - `+ ( targetCurvature - curCurvature ) * dt * speed_normalized`
        - `* speed_normalized` because if car dosen't move,  the curvature should not change

- How to add player's control 
    - when in a curve track,  the horizontal offset of the car is accumulated, and player must control the car to cancel out that offset to keep the car moving inside the track.


