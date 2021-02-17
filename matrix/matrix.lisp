; A list is a 1-D array of numbers.
; A matrix is a 2-D array of numbers, stored in row-major order.

; If needed, you may define helper functions here.

; AreAdjacent returns true iff a and b are adjacent in lst.
(defun are-adjacent (lst a b)
    (cond ( (null (cdr lst) ) nil)
            ( (and (eql (car lst) a) (eql (car (cdr lst) ) b) ) t)
            ( (and (eql (car lst) b) (eql (car (cdr lst) ) a) ) t)
            (t (are-adjacent(cdr lst) a b ))
    )
)

; Transpose returns the transpose of the 2D matrix mat.
(defun transpose (matrix)
    (if (null matrix) nil
        (apply #'mapcar #'list matrix)
    )
)

; AreNeighbors returns true iff a and b are neighbors in the 2D
; matrix mat.
(defun are-neighbors (matrix a b)
    (cond ( (helper-neighbors matrix a b) T)
          ( (helper-neighbors (transpose matrix) a b ) T)
          ( T nil )
    )
)

; Helper function to use recursion in are-neighbors
(defun helper-neighbors (matrix a b)
    (cond ( (null matrix) nil)
          ( (are-adjacent (car matrix) a b) t)
          (t (are-neighbors (cdr matrix) a b) )    
    )
)
