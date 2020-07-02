import numpy as np


def normalize( v ):
    norm = np.linalg.norm(v)
    if norm == 0 :
        return v
    else:
        return v/norm


def newTranslatesMatrix( x,y,z ):
    mat = np.identity(4)
    mat[0][3] = x
    mat[1][3] = y
    mat[2][3] = z
    return mat


# https://lmb.informatik.uni-freiburg.de/people/reisert/opengl/doc/glTranslate.html
def newLookAtMatOnTheFly( eyePosition3D, center3D , upVector3D ):
    forward = np.zeros(3)
    forward[0] = center3D[0] - eyePosition3D[0]
    forward[1] = center3D[1] - eyePosition3D[1]
    forward[2] = center3D[2] - eyePosition3D[2]
    forward = normalize( forward )

    # Side = forward x up
    side = normalize(np.cross( forward, upVector3D ))
    # Recompute up as: up = side x forward
    # recalculate up
    up = np.cross( side, forward )  # should already be normalized
    print ( upVector3D, up  )

    m16 = np.zeros(16)  # column based
    m16[0] = side[0]
    m16[4] = side[1]
    m16[8] = side[2]
    m16[12] = 0.0
    # --------------------
    m16[1] = up[0]
    m16[5] = up[1]
    m16[9] = up[2]
    m16[13] = 0.0
    # --------------------
    m16[2] = -forward[0]
    m16[6] = -forward[1]
    m16[10] = -forward[2]
    m16[14] = 0.0
    # --------------------
    m16[3] = m16[7] = m16[11] = 0.0
    m16[15] = 1.0

    # reshape row based, so need .T
    mat = m16.reshape( 4,4 ).T

    # glTranslated(-eyex, -eyey, -eyez)
    matTranslate = newTranslatesMatrix( -eyePosition3D[0],-eyePosition3D[2],-eyePosition3D[2] )
    return np.matmul( mat, matTranslate )


# not efficient , it is used for verification
def newLookAtMatOnTheFly2( eyePosition3D, center3D , upVector3D ):
    forward = np.zeros(3)
    # for camera , forward is mapped to -Z
    forward[0] = -(center3D[0] - eyePosition3D[0])
    forward[1] = -(center3D[1] - eyePosition3D[1])
    forward[2] = -(center3D[2] - eyePosition3D[2])
    forward = normalize( forward )

    # right hand , upxforward
    side = normalize(np.cross( upVector3D, forward ))
    up = np.cross( forward, side )

    # camera roated matrix
    m16 = np.zeros(16)  # column based
    for i in range(3):
        m16[i+0] = side[i]  # col 0 , 0-3
        m16[i+4] = up[i]  # col 1  ,  4-7
        m16[i+8] = forward[i]  # col 2  , 8-11
    m16[15] = 1

    matRot = m16.reshape(4,4).T
    matTranslate = newTranslatesMatrix( eyePosition3D[0],eyePosition3D[1],eyePosition3D[2] )
    # rotate first, then translate  
    matCamera = np.matmul( matTranslate, matRot  ) 
    # print ( "camera mat:" )
    # print( matCamera )
    return np.linalg.inv( matCamera )



if __name__ == '__main__':

    print ( "trans: 1,2,3 ============================")
    t = newTranslatesMatrix( 1,2,3 )
    print( t )
    print ( "inv trans:")
    print( np.linalg.inv( t ) )
    print ( np.matmul(t, np.linalg.inv(t)))

    r = np.array( [ normalize( np.array( [1,2,3,0] ) ), normalize( np.array( [4,4,4,0] ) ), normalize( np.array( [7,7,9,0] ) ), [0,0,0,1]  ]  )
    r = r.reshape(4,4).T
    print ("rot: ==============================")
    print (r)
    print ("inv rot:")
    print (np.linalg.inv(r))
    print ( np.matmul( r, np.linalg.inv(r) ) )

    print ( "r * t ============================" )
    print ( np.matmul( r, t ) )
    print ( "t * r " )
    print ( np.matmul( t, r ) )

    print ("viewer ==============================")
    eye = np.array([1,2,3])
    center = np.array([ 4,3,4])
    up = np.array( [0,1,0] )
    print (  newLookAtMatOnTheFly( eye,center,up ) )
    print ( "unnormalized up" )
    up = np.array( [0,2,0] )
    print (  newLookAtMatOnTheFly( eye,center,up ) )
    print ( "inv viewer matrix" )
    print ( np.linalg.inv( newLookAtMatOnTheFly( eye,center,up ) ) )

    print ("viewer 2 ==============================")
    print (  newLookAtMatOnTheFly2( eye,center,up ) )

    print ( "viewer * [11,2,6,1]=" )
    print ( np.matmul( newLookAtMatOnTheFly2( eye,center,up ) , np.array( [11,2,6,1] )  ) )

    ma = np.array( [ 1,0,0,0, 2,1,0,0, 3,4,5,6, 0,1,0,1 ] ).reshape( 4,4 ).T
    mb = np.array( [ 1,2,3,4, 1,1,0,0, 7,6,5,6, 1,1,0,1 ] ).reshape( 4,4 ).T
    print( "ma * mb" )
    print ( np.matmul( ma,mb ) )



