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

    verts := make(  []Vec3D , 128  )
    faces := make(  []Triangle , 128  )

    var v Vec3D
    var v_idx [3]int

    cnt := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        cnt ++
        line := scanner.Text()
        if len(line) == 0 {
            continue
        }
        if line[:1] == "v" {
            fmt.Sscanf( line[1:], "%f %f %f", &v.X, &v.Y, &v.Z )
            verts = append( verts, v )
            // log.Println( cnt, v.X, v.Y, v.Z )
        } else if line[:1] == "f" {
            
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}


