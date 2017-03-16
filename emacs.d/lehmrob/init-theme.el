;; Init for the theme
;; (maybe-require-package 'minimal-theme)
(when (display-graphic-p)
  (progn
	(load-theme 'adwaita t))
  (progn
	(custom-set-variables '(frame-background-mode 'light))))
(provide 'init-theme)
