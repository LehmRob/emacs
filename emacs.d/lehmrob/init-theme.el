;; Init for the theme
(when (display-graphic-p)
  (progn
	(maybe-require-package 'minimal-theme)
	(load-theme 'minimal-light t))
  (progn
	(custom-set-variables '(frame-background-mode 'light))))
(provide 'init-theme)
