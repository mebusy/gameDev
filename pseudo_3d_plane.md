
# Pseudo 3D Plane

This was a technique made popular by the super nintendo.  I'm going to use a super mario kart as an example where an image is is taken as the ground plane and it is roated and scaled and translated around the cameras necessary to give the impression of a 3D environment. This was known as mode 7 which for the super Nintendo was just really implementing an affine transform in hardware. 

I'm not going to use affine transforms directly because I want to play with some of the parameters and see what happens.

I'm going to create a big sprite and we know that behave just like textures. We can sample texture from 0.0 to 1.0.  This makes sampling the texture scale invariant. 

I'm going to treat this big picture as if it were my world map , or the mario kart track in this case.  In this map, I'm going to have a postion which represent where are the camera is and I'm going to call this position "world X,Y".  The camera also have an angle associated with it, the angle is what direction is the camera looking at in this plane. 

![](imgs/mode7_0.png)

From this information, I'm going to create a viewing frustum but I need a little bit more information first. 

Firstly I need a field of view, call that θ.


