(maybe-require-package 'avy)
(avy-setup-default)

;; Add some special keybindings
(global-set-key (kbd "C-,") 'avy-goto-char)
(global-set-key (kbd "C-'") 'avy-goto-char-2)
(global-set-key (kbd "C-c C-f") 'avy-goto-line)
(global-set-key (kbd "C-c C-w") 'avy-goto-word-1)
(global-set-key (kbd "C-c C-e") 'avy-goto-word-0)

(provide 'init-avy)
