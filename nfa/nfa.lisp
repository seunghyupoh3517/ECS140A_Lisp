;; You may define helper functions here
(defun reachable (transition start final input) 
    ;; TODO: Incomplete function
    ;; The next line should not be in your solution.
    (cond ((null input)
        ;; when input is null, check whether the start and final is equal
            (funcall (lambda (s1 s2) (eq s1 s2)) start final )) 

        ;; when input is not null, proceed to the next state with car(input)
        ;; then recursive call with next state and cdr(input)
        ;; Multiple possible next states can be called interatively 
        ;; transition f returns list of possible states
          ((not(null input))
            (setq next (funcall transition start (car input))) ;; list of next states 
        ;; Prolbem#1 - how to iteratively call the reachable() with different next starting states
            (reach_helper1 transition next final (cdr input)))
    ) 
) 

;; Send the list of next states which is generated with start state and input
(defun reach_helper1 (transition nextlist final input)
    (cond ((null input)
        ;; Problem#2 - comparing the element of list and atom is not possible, mapcar takes arg list only
        ;; i.e. nextlist (0 1 3) -> truelist (nil nil t) when final is 3
        ;; if any of the possible next state is final state, reachable
        ;; (reach_helper2 truelist))) 
            (reach_helper2 (setq truelist (mapcar #'(lambda (s1) (eq s1 final)) nextlist))))
        
          ((not(null input))
            (or (reachable transition (car nextlist) final input)
            (reach_helper3 transition (cdr nextlist) final input))
          )
    )
)

;; truelist from reach_helper1 can be checked if it reaches the final state here
(defun reach_helper2 (start)
    (if (not (null start))
            (cond 
                ((eq (car start) t) (eval t))
                ((eq (car start) nil) (reach_helper2 (cdr start)))
            )
    )
)

;; In order to iteratively go all the state in the list of next states
(defun reach_helper3 (transition nextlist final input)
    (if (not (null (car nextlist))) 
        (or (reachable transition (car nextlist) final input)
        (reach_helper3 transition (cdr nextlist) final input)))
)
