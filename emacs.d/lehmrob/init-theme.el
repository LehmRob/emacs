;; Init for the theme
(when (display-graphic-p)
  (progn
	(maybe-require-package 'leuven-theme)
	(load-theme 'leuven t))
  (progn
  	(custom-set-variables '(frame-background-mode 'light))))
(provide 'init-theme)
