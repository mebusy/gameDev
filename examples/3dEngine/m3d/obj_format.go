package m3d

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func (self *Mesh) LoadFromObj( filename string ) {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    verts := make(  []Vec3D ,0, 64  )
    faces := make(  []Triangle ,0, 64  )

    var f [3]int

    cnt := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        cnt ++
        line := scanner.Text()
        if len(line) == 0 {
            continue
        }
        if line[:1] == "v" {
            var v Vec3D
            fmt.Sscanf( line[1:], "%f %f %f", &v.X, &v.Y, &v.Z )
            verts = append( verts, v )
            // log.Println( cnt, v.X, v.Y, v.Z )
        } else if line[:1] == "f" {
            fmt.Sscanf( line[1:], "%d %d %d", &f[0], &f[1], &f[2] )
            // log.Println( cnt, f[0], f[1], f[2], verts[f[0]-1] )
            // vertex index starts from 1
            faces = append(faces,
                Triangle{ [3]Vec3D{verts[f[0]-1], verts[f[1]-1],verts[f[2]-1]}, 0 } )
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    // log.Println( verts )
    self.Tris = faces
}


