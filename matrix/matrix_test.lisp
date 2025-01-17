;;;;;;;;;;;;;;;;;;;;;;;;
;;; pre-testing prep ;;;
;;;;;;;;;;;;;;;;;;;;;;;;

(load "../lisp-unit.lisp")

(use-package :lisp-unit)

(load "matrix.lisp")

(remove-tests :all)

(setq *print-failures* t)

;;;;;;;;;;;;;;;;;;;;;;;;
;;; test definitions ;;;
;;;;;;;;;;;;;;;;;;;;;;;;

(define-test test-are-adjacent
    (assert-equal NIL (are-adjacent NIL 1 1))
    (assert-equal NIL (are-adjacent '() 1 1))
    (assert-equal NIL (are-adjacent '(1) 1 1))
    (assert-equal T (are-adjacent '(1 1) 1 1))
    (assert-equal NIL (are-adjacent '(1 2 3) 1 3))
    (assert-equal NIL (are-adjacent '(1 2 3) 3 1))
    (assert-equal T (are-adjacent '(1 2 3) 1 2))
    (assert-equal T (are-adjacent '(1 2 3) 2 3))
    (assert-equal T (are-adjacent '(1 2 3) 3 2))
    (assert-equal T (are-adjacent '(1 2 3) 2 1))
    (assert-equal NIL (are-adjacent '(1 2 1) 1 1))
    (assert-equal T (are-adjacent '(1 2 1) 1 2))
)

(define-test test-transpose
    (assert-equal () (transpose ()))
    (assert-equal '( (1) ) (transpose '( (1) )))
    (assert-equal '( (1) (2) (3) (4) ) (transpose '( (1 2 3 4) )))
    (assert-equal '( (1 2 3 4) ) (transpose '( (1) (2) (3) (4) )))
    (assert-equal '( (1 2) (3 4) ) (transpose '( (1 3) (2 4) )))
    (assert-equal '( (1 3) (2 4) ) (transpose '( (1 2) (3 4) )))
)

(define-test test-are-neighbors
    (assert-equal NIL (are-neighbors () 1 2))
    (assert-equal T   (are-neighbors '( (1 2 3) ) 1 2))
    (assert-equal NIL (are-neighbors '( (1 2 3) ) 1 3))
    (assert-equal T   (are-neighbors '( (1) (2) (3) ) 1 2))
    (assert-equal T   (are-neighbors '( (1 2 3) (4 5 6) ) 1 2))
    (assert-equal T (are-neighbors '( (1 2 3) (4 5 6) ) 2 3))
    (assert-equal T (are-neighbors '( (1 2 3) (4 5 6) ) 2 5))
    (assert-equal T (are-neighbors '( (1 2 3) (4 5 6) ) 3 6))
    (assert-equal T (are-neighbors '( (1 2 3) (4 5 6) ) 1 4))
    (assert-equal NIL (are-neighbors '( (1 2 3) (4 5 6) ) 2 6))
    (assert-equal NIL (are-neighbors '( (1 2 3) (4 5 6) ) 1 6))
    (assert-equal T (are-neighbors '( (1 2 3 4) (5 6 7 8) (9 10 11 12) (13 14 15 16) ) 10 6))
    (assert-equal T (are-neighbors '( (1 2 3 4) (5 6 7 8) (9 10 11 12) (13 14 15 16) ) 12 16))
    (assert-equal NIL (are-neighbors '( (1 2 3 4) (5 6 7 8) (9 10 11 12) (13 14 15 16) ) 7 15))
    (assert-equal NIL (are-neighbors '( (1 2 3 4) (5 6 7 8) (9 10 11 12) (13 14 15 16) ) 9 4))
)

;;;;;;;;;;;;;;;;;
;;; run tests ;;;
;;;;;;;;;;;;;;;;;

(run-tests :all)
