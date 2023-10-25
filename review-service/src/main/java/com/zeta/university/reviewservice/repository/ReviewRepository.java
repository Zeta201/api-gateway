package com.zeta.university.reviewservice.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.zeta.university.reviewservice.entity.Review;

public interface ReviewRepository extends JpaRepository<Review, Long>{

}
