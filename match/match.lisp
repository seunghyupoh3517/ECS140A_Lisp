; You may define helper functions here

;; (defun matchsame (ls1 ls2)
;;   (equal ls1 ls2) 
;; )

;; (defun matchQ (ls1 ls2) ;detect ? keep going
;;   (equal (car ls1) '?) (cdr ls1)(cdr ls2)
;; )

;; (defun matchE (ls1 ls2)
;;   (cond
;;     ((equal (car ls1 ) '!) (matchE(cdr ls1) (cdr ls2)))
;;     ((equal (car ls1) (car ls2) T)

;;   )
;; ))

(defun match (pattern assertion)
  ;; TODO: incomplete function. 
  ;; The next line should not be in your solution.
  (cond
    ; (equal (pattern assertion null) t)
    ((and (null pattern) (null assertion) T)) ;check if empty

    ((equal pattern assertion) T) ; all same return T and end
    ((and (null pattern) (not (null assertion))) nil )
    ((and (null assertion) (not (null pattern))) nil )


    ((equal (car pattern) '?) (match(cdr pattern) (cdr assertion))) ;detect ?

    ((and (equal (car pattern) '!) (equal (cdr pattern) nil) T)) 
    ; see ! at the end of pattern
    ((equal (car pattern) '!) (or (match pattern (cdr assertion)) (match(cdr pattern) (cdr assertion))) )
    ; match pattern will hold the position of pattern and iterate through assertion
    ; if it reach the end of assertion, it will go to line 28 and return nil
    ; go back one step back with pattern and assertion, so the assertion will not be nil
    
    ; (!) (a b c) 
    ; (!) (b c)
    ; (!) (c)
    ; (!) nil
    ; go back
    ; (!) (c)
    ; nil nil
    ; T

    ; one "or" fail, could go to another.
    ; first "or" could find if there is anything after ! that match(go all the way to the end)
    ; see nil go back one step
    ; second "or" if there is just one ! means one letter(impoetant condition for moving forward)


    ((equal (car pattern) (car assertion))  (match(cdr pattern) (cdr assertion)) );rec iterate through
    ; if we miss all matches from cond above but still match the first letter, moving forward
    ; first letter same, moving forward
  )
)
