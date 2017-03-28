;; Init for the theme
(maybe-require-package 'basic-theme)
(when (display-graphic-p)
  (progn
	(load-theme 'basic t))
  (progn
	(custom-set-variables '(frame-background-mode 'light))))
(provide 'init-theme)
